#!/bin/bash

./get_kubeconfig/basis_1_main > kubeconfig
export KUBECONFIG=$(pwd)/kubeconfig

helm repo add jenkins https://charts.jenkins.io
helm upgrade jenkins jenkins/jenkins -n jenkins --create-namespace --install --values ./helm/jenkins.yml

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm upgrade ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx --create-namespace --install
