#!/bin/bash

bash_path=$(cd "$(dirname "$0")";pwd)
source $bash_path/base.config

if [[ "$(whoami)" != "root" ]]; then
        echo "please run this script as root and on master node." >&2
        exit 1
fi

echo -e "---------------------\033[34m start deploy microservice to k8s \033[0m---------------------"
echo -e "-------------\033[34m Checking Basic Env Status \033[0m------------"
echo "Checking Harbor health..."
# 0. Basic Preparation: Harbor(we use Harbor instead of Dockerhub), K8S, Docker, SSH Keys already exist.
# 0.1. check Harbor is existed
# shellcheck disable=SC2154
harbor_domain=${harbor_ip}
response=$(curl -s -o /dev/null -w "%{http_code}" http://${harbor_domain}/api/v2.0/health)
if [ "$response" = "200" ]; then
    echo "Harbor server is accessible."
else
    echo "Harbor server at http://${harbor_domain} not accessible or Health API not responding with HTTP 200." >&2
    exit 1
fi
# 0.2. check K8S is existed and healthy
# 0.2 检查Kubernetes集群是否存在且健康
echo "Checking Kubernetes cluster health..."
kubectl get nodes
if [ $? -ne 0 ]; then
    echo "Failed to get Kubernetes nodes, please check your K8S setup." >&2
    exit 1
fi

kubectl get pods --all-namespaces
if [ $? -ne 0 ]; then
    echo "Failed to get pods, please check your K8S cluster health." >&2
    exit 1
fi
# 0.3. check Docker is existed
echo "Checking Docker..."
docker --version
if [ $? -ne 0 ]; then
    echo "Docker is not installed or not running." >&2
    exit 1
fi
# 0.4. check SSH could be used
echo "Checking SSH..."
# 尝试SSH到localhost来确认SSH服务正常
ssh localhost 'exit'
if [ $? -ne 0 ]; then
    echo "SSH is not working properly." >&2
    exit 1
fi

echo "-------------\033[32m Passed all checks, ready to deploy microservice. \033[0m------------"
# 1. Make sure Harbor & K8S can use http protocol
# 1.1. Modify Harbor ???
# 1.2. Add following string into /etc/docker/daemon.json to enable http protocol
#sudo cp /etc/docker/daemon.json /etc/docker/daemon.json.bak  # 备份
#sudo sed -i '$s/}$/},/' /etc/docker/daemon.json  # 设置好末尾逗号
#sudo sed -i '$i\  "insecure-registries": ["192.168.2.14:214"]' /etc/docker/daemon.json  # 插入语句
# 1.3. Restart Docker to make the changes take effect
#sudo systemctl restart docker
echo -e "-------------\033[34m Docker Pushing Harbor \033[0m------------"
# Docker Push
docker build -t 192.168.2.14:214/kratos_test/test_server1:v1.0 .  # （需要在docker所在的机器上执行）构建镜像
docker login 192.168.2.14:214 -u admin -p Ingruharbor123456  # （需要在docker所在的机器上执行）如果要用内网Harbor，只能登录，否则默认推送到dockerhub
docker push 192.168.2.14:214/kratos_test/test_server1:v1.0  # （需要在docker所在的机器上执行）推送镜像
echo -e "-------------\033[32m Harbor Archived \033[0m------------"
# 2. Modify ./build/deployment.yaml to use the correct image and name the app correctly

# 2.1 Create k8s-docker-secret for K8S pulling images from Harbor, and replace name at 'imagePullSecrets'

#"
#kubectl create secret docker-registry myregistrykey \
#--docker-server=192.168.2.14:214 \
#--docker-username=root \
#--docker-password=123456
#"

# 2.2 Modify rest configurations in deployment.yaml
# 3. Create namespace(if it doesn't exist)
# 3.1. Create New namespace
echo -e "-------------\033[34m K8S Cleaning Necessary \033[0m------------"
kubectl create namespace ${namespace} 2>/dev/null
if kubectl get namespace ${namespace}; then
    echo "Namespace '${namespace}' is successfully created or already exists."
else
    echo "Failed to create namespace '${namespace}'."
fi
# 4. Delete former app's cache on each node of K8S
# FIXME it's too dangerous to delete pod directly
# 遍历每个 IP 地址
for ip in "${k8s_ips[@]}"
do
    echo "Connecting to $ip to clean docker caches for ${pod_name}..."
    # SSH 进入每一台机器
    ssh root@$ip << EOF
        # 查找匹配 TARGET 的容器，并确认是否删除
        echo "Found containers matching '${pod_name}':"
        docker ps -a | grep '${pod_name}'
        echo "Do you want to stop and remove these containers? (y/N)"
        read -r answer
        if [[ \$answer =~ ^[Yy]$ ]]
        then
            docker ps -a | grep '${pod_name}' | awk '{print \$1}' | xargs -r docker stop | xargs -r docker rm
            echo "Containers matching '${pod_name}' stopped and removed."
        else
            echo "Containers removal skipped."
        fi

        # 查找匹配 TARGET 的镜像，并确认是否删除
        echo "Found images containing '${pod_name}':"
        docker images | grep '${pod_name}'
        echo "Do you want to remove these images? (y/N)"
        read -r answer
        if [[ \$answer =~ ^[Yy]$ ]]
        then
            docker images | grep '${pod_name}' | awk '{print \$3}' | xargs -r docker rmi
            echo "Images containing '${pod_name}' removed."
        else
            echo "Images removal skipped."
        fi
EOF
    echo "Cleaned on $ip."
done

echo "-------------\033[32m Related caches are cleaned carefully for ${pod_name}. \033[0m------------"

# 5. Deploy deployment.yaml to K8S cluster
# 5.1. Delete running related pod with specific name
echo -e "-------------\033[34m Deleting Related Pods \033[0m------------"
kubectl get pods -n ${namespace} | grep ${pod_name} | awk '{print $1}' | xargs kubectl delete pod
echo "-------------\033[32m Related pods are cleaned carefully for ${pod_name}. \033[0m------------"
# 5.2. Execute command on the control-panel
echo -e "-------------\033[34m Deploying deployment.yaml \033[0m------------"
kubectl apply -f deployment.yaml
# 检查操作的返回状态
if [ $? -eq 0 ]; then
    # 如果kubectl命令成功执行， $? 将返回 0
    echo "Deployment applied successfully."
    # 检查deployment的状态，确保所有的Pod都成功启动了
    kubectl rollout status deployment/${pod_name} -n ${namespace}
    if [ $? -eq 0 ]; then
        echo "Deployment rollout succeeded."
    else
        echo "Deployment rollout failed."
    fi
else
    # 如果kubectl命令运行失败， $? 将返回非0值
    echo "Failed to apply deployment."
fi
echo "-------------\033[32m Deployed deployment.yaml for ${pod_name}. \033[0m------------"
# 5.3. Get callback for checking the status of the app
# 6. Expose the service to the internet
echo -e "-------------\033[34m Reexpose Service for ${pod_name}. \033[0m------------"
# 6.1. Check status of deployments and services
kubectl get deployments --all-namespaces
kubectl get services --all-namespaces
# 6.2. Expose the service to the internet
kubectl apply -f service.yaml
if [ $? -eq 0 ]; then
   # 如果kubectl命令成功执行， $? 将返回 0
   echo "Deployment applied successfully."
else
   # 如果kubectl命令运行失败， $? 将返回非0值
   echo "Failed to apply deployment."
fi
# 6.3. Check the status of the service
# 7. Check the status of the app


