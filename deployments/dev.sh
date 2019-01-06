helm install --name unicampus-postgres --set postgresqlPassword=postgres --set replication.slaveReplicas=0 stable/postgresql
helm install --name unicampus-elasticsearch --set master.replicas=0 --set client.replicas=0 stable/elasticsearch
kubectl apply -f ./dev.yml
