package testing

import "github.com/chnsz/golangsdk/openstack/cce/v3/partitions"

const Output = `{
    "kind":"Partition",
    "apiVersion":"v3",
    "metadata":{
        "name":"cn-south-1-ies-fstxz"
    },
    "spec":{
        "hostNetwork":{
            "subnetID":"d7131ed5-f813-4dbc-86f8-bcbdc07dce6f"
        },
        "containerNetwork":[
            {
                "subnetID":"b2f23c46-edaa-4e66-b82f-50edafa638f5"
            },
            {
                "subnetID":"dee746d5-6c78-43fb-bc36-ac26c581a3ec"
            }
        ],
        "publicBorderGroup":"cn-south-1-ies-fstxz",
        "category":"IES"
    }
}`

var Expected = &partitions.Partitions{
	Kind:       "Partition",
	Apiversion: "v3",
	Metadata:   partitions.Metadata{Name: "cn-south-1-ies-fstxz"},
	Spec: partitions.Spec{
		Category:          "IES",
		PublicBorderGroup: "cn-south-1-ies-fstxz",
		HostNetwork: partitions.HostNetwork{
			SubnetID: "d7131ed5-f813-4dbc-86f8-bcbdc07dce6f",
		},
		ContainerNetwork: []partitions.ContainerNetwork{
			{
				SubnetID: "b2f23c46-edaa-4e66-b82f-50edafa638f5",
			},
			{
				SubnetID: "dee746d5-6c78-43fb-bc36-ac26c581a3ec",
			},
		},
	},
}
