sudo docker build -t goose:blockexplorer -f Dockerfile.goose .
sudo docker build -t sync:blockexplorer -f Dockerfile.sync .
sudo docker build -t rest:blockexplorer -f Dockerfile.rest .
sudo docker tag goose:blockexplorer 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-goose
sudo docker tag sync:blockexplorer 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-sync
sudo docker tag rest:blockexplorer 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-rest
sudo docker push 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-goose
sudo docker push 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-sync
sudo docker push 086918792671.dkr.ecr.us-east-1.amazonaws.com/blockexplorer-rest
