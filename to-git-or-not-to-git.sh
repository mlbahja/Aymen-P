curl --GET "https://learn.zone01oujda.ma/assets/superhero/all.json" | jq '.[] | select(.id ==170) | {name: .name, power: .powerstats.power, gender: .appearance.gender}'
