import yaml

with open('machine.yaml', 'r') as machines:
    nodes = yaml.safe_load(machines)
print(nodes['machines'][0]['pd_servers'])
with open('topology.yaml', 'r') as file:
    tikv_config = yaml.safe_load(file)
tikv_config['pd_servers'][0]['host']= nodes['machines'][0]['pd_servers']
tikv_config['tikv_servers'][0]['host']= nodes['machines'][1]['kv_1']
tikv_config['tikv_servers'][1]['host']= nodes['machines'][2]['kv_2']
tikv_config['tikv_servers'][2]['host']= nodes['machines'][3]['kv_3']
tikv_config['monitoring_servers'][0]['host']= nodes['machines'][4]['monitor']
tikv_config['grafana_servers'][0]['host']= nodes['machines'][4]['monitor']
tikv_config['alertmanager_servers'][0]['host']= nodes['machines'][4]['monitor']

print(tikv_config['tikv_servers'])
print(tikv_config)
with open('topology.yaml', 'w') as file:
    yaml.dump(tikv_config, file, sort_keys=False)