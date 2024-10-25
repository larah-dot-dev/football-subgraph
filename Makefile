deploy:
	caprover deploy -h https://football-subgraph.larah.dev -b master -a football-subgraph --default

fmt:
	find . -type f -name "*.go" | go fmt
