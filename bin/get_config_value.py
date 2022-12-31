import json, os, sys
gopath = os.getenv("goCoreAppPath")
base = gopath + "/bin/globalcache/"
config = json.load(open(gopath + '/webConfig.json', 'r'))
ptrOne = sys.argv[1]
ptrTwo = sys.argv[2]
print config[ptrOne][ptrTwo]
