# Installation of Tenant Operations Controller alies deployer

## Kubernetes Cluster

### Pre-requisites:

1. Mongodb connection URL
2. Mongodb username and password
3. AWS_ROLE_ARN - "TenantOperationsController-eks"
4. AWS_NODE_ARN - "TenantOperationsController-ec2" 
5. AWS_SECURITY_GROUPS - IDs for security groups
6. AWS_SUBNETS - IDs of subnets in different AZs. minimum subnet requirement is 2.
   example of AZs: us-east-1a,us-east-1b,us-east-1c,us-east-1d.
#### 7. Important the user which was created in the README.md of goapp should be provided in the secret.yaml file.

You can install by running the following command from the home directory

```
sudo kubectl apply -f install/ -n <namespace>
```

