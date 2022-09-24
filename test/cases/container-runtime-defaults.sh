#!/usr/bin/env bash
set -euo pipefail

exit_code=0
TEMP_DIR=$(mktemp -d)

echo "--> Should allow dockerd as container runtime when below k8s version 1.24"
KUBELET_VERSION="v1.20.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    --container-runtime dockerd \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi

echo "--> Should allow containerd as container runtime when below k8s version 1.24"
KUBELET_VERSION="v1.20.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    --container-runtime containerd \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi

echo "--> Should have default container runtime when below k8s version 1.24"
KUBELET_VERSION="v1.20.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi

echo "--> Should not allow dockerd as container runtime when at or above k8s version 1.24"
export KUBELET_VERSION="v1.24.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    --container-runtime dockerd \
    test || exit_code=$?

echo "EXIT CODE $exit_code"
if [[ ${exit_code} -eq 0 ]]; then
    echo "❌ Test Failed: expected a non-zero exit code but got '${exit_code}'"
    exit 1
fi
exit_code=0

echo "--> Should allow containerd as container runtime when at or above k8s version 1.24"
KUBELET_VERSION="v1.24.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    --container-runtime containerd \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi

echo "--> Should have default container runtime when at or above k8s version 1.24"
KUBELET_VERSION="v1.24.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi

echo "--> Should ignore docker-specific flags when at or above k8s version 1.24"
KUBELET_VERSION="v1.24.15-eks-ba74326"
run ${TEMP_DIR} /etc/eks/bootstrap.sh \
    --b64-cluster-ca dGVzdA== \
    --apiserver-endpoint http://my-api-endpoint \
    --enable-docker-bridge true \
    --docker-config-json "{\"some\":\"json\"}" \
    test || exit_code=$?

if [[ ${exit_code} -ne 0 ]]; then
    echo "❌ Test Failed: expected a zero exit code but got '${exit_code}'"
    exit 1
fi
