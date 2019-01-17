kubectl port-forward --namespace default svc/unicampus-postgresql 5432:5432 &
kubectl port-forward --namespace default svc/unicampus-elasticsearch-client 9200:9200 &
kubectl port-forward --namespace default svc/unicampus-neo4j 7474:7474 &
