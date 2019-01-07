helm install --name unicampus-postgres --set postgresqlPassword=postgres --set replication.slaveReplicas=0 stable/postgresql
helm install --name unicampus-elasticsearch --set master.replicas=2 --set client.replicas=0 stable/elasticsearch
# helm install --name unicampus-neo4j --set acceptLicenseAgreement=yes --set authEnabled=true --set neo4jPassword=neosecret --set core.numberOfServers=1 --set readReplica.numberOfServers=0 stable/neo4j
kubectl apply -f ./dev.yml
