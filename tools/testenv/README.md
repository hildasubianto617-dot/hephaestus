# testenv
Library for creating Kubernetes environments

## Average Convergence Times

| Cluster Type | Creation | Destruction |
|--------------|----------|-------------|
| GKE          | ~ 7m     | ~ 9m58s     |
| EKS          | ~        | ~           |
| AKS          | ~ 4m58s  | ~ 5m41s     |

## GKE prerequisites

The GKE environment does not create service accounts or change IAM policy, so
the identity that runs it needs no IAM administration rights. Create these once
per project and pass their emails to `GKEConfig`.

```bash
PROJECT_ID=<project>

# Service account the controller impersonates to push to Artifact Registry.
gcloud iam service-accounts create testenv-gke-agent --project="$PROJECT_ID" \
  --display-name="testenv GKE agent"
gcloud projects add-iam-policy-binding "$PROJECT_ID" \
  --member="serviceAccount:testenv-gke-agent@${PROJECT_ID}.iam.gserviceaccount.com" \
  --role="roles/artifactregistry.writer"
# Workload Identity: the controller's Kubernetes service account (default/hephaestus)
# in any cluster in this project may act as the agent account.
gcloud iam service-accounts add-iam-policy-binding \
  "testenv-gke-agent@${PROJECT_ID}.iam.gserviceaccount.com" --project="$PROJECT_ID" \
  --role="roles/iam.workloadIdentityUser" \
  --member="serviceAccount:${PROJECT_ID}.svc.id.goog[default/hephaestus]"

# Service account the cluster nodes run as.
gcloud iam service-accounts create testenv-gke-node --project="$PROJECT_ID" \
  --display-name="testenv GKE node"
for role in roles/container.defaultNodeServiceAccount roles/artifactregistry.reader; do
  gcloud projects add-iam-policy-binding "$PROJECT_ID" \
    --member="serviceAccount:testenv-gke-node@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role="$role"
done
```

The identity that runs Terraform needs, in that project:

- `roles/container.admin`, `roles/compute.networkAdmin` and `roles/artifactregistry.admin`
- `roles/iam.serviceAccountUser` on the node service account only, to attach it
  to the nodes:

```bash
gcloud iam service-accounts add-iam-policy-binding \
  "testenv-gke-node@${PROJECT_ID}.iam.gserviceaccount.com" --project="$PROJECT_ID" \
  --role="roles/iam.serviceAccountUser" \
  --member="serviceAccount:<terraform-identity-email>"
```
