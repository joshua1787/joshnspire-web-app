# 🌟 JoshNspire Web App - DevOps Project

This project demonstrates a **production-ready** web application deployed using **Kubernetes (EKS)**, **Helm**, **ArgoCD**, and **GitHub Actions CI/CD**.

---

## 🚀 Project Overview

- **Frontend:** Simple static HTML website (DevOps Courses Platform)
- **Backend:** GoLang server
- **Containerized:** Docker
- **Deployed on:** AWS Elastic Kubernetes Service (EKS)
- **Load Balancer:** AWS ALB + Ingress Controller (NGINX)
- **CI/CD:** GitHub Actions (build, test, lint, Docker push, Helm update)
- **GitOps:** ArgoCD for auto-deployment of new versions

---

## 🛠️ Technologies Used

- **AWS EKS** - Kubernetes Cluster
- **Helm** - Application Packaging and Deployment
- **ArgoCD** - GitOps Continuous Deployment
- **GitHub Actions** - CI/CD Pipeline
- **Docker** - Containerization
- **NGINX Ingress Controller** - Load Balancing and Routing
- **GoLang** - Web server
- **HTML/CSS** - Static Website

---

## 📂 Project Structure

- `.github/workflows/ci.yaml` — GitHub Actions workflow file
- `helm/joshnspire-web-app-chart/` — Helm chart for Kubernetes deployment
- `k8s/` — Kubernetes manifests (optional)
- `static/` — Static HTML pages for frontend
- `main.go` — GoLang server code
- `main_test.go` — GoLang unit test file
- `Dockerfile` — Docker image build instructions
- `go.mod` — Go dependencies file (module setup)

## 🔥 Features

- ✅ Automatic Docker image build and push to DockerHub
- ✅ Helm chart tag updates automatically after every build
- ✅ Continuous Deployment using ArgoCD GitOps
- ✅ Ingress Controller Load Balancer setup with DNS routing
- ✅ Production-grade Kubernetes deployment following best practices

---

## 📸 Screenshots

_(You can upload screenshots here later if you want)_

---

## 🙌 Author

**Joshua Veeraiah**  
🔗 [LinkedIn Profile](https://www.linkedin.com/in/joshua1787/)  
🏅 AWS Certified DevOps Engineer - Professional

---

## 📢 Important

> This project is for **learning**, **portfolio building**, and **showcasing** real-world DevOps project skills to recruiters and companies.


