curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq "[map(select(.id ==$HERO_ID))] | .[] | .[] .connections.relatives" | tr -d '"' 
