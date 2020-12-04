package main

import (
	"fmt"
	"io/ioutil"
	"os"

	cblog "github.com/cloud-barista/cb-log"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	osdrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/openstack"
	_ "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/openstack/connect"
	_ "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/drivers/openstack/resources"
	idrv "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces"
	irs "github.com/cloud-barista/cb-spider/cloud-control-manager/cloud-driver/interfaces/resources"
)

var cblogger *logrus.Logger

func init() {
	// cblog is a global variable.
	cblogger = cblog.GetLogger("CB-SPIDER")
}

func testImageHandler(config Config) {
	resourceHandler, err := getResourceHandler("image")
	if err != nil {
		panic(err)
	}

	imageHandler := resourceHandler.(irs.ImageHandler)

	fmt.Println("Test ImageHandler")
	fmt.Println("1. ListImage()")
	fmt.Println("2. GetImage()")
	fmt.Println("3. CreateImage()")
	fmt.Println("4. DeleteImage()")
	fmt.Println("5. Exit")

	imageId := irs.IID{
		NameId: "Fedora27-k8s", // Ubuntu 16.04
	}

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListImage() ...")
				if list, err := imageHandler.ListImage(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(list)
				}
				fmt.Println("Finish ListImage()")
			case 2:
				fmt.Println("Start GetImage() ...")
				if imageInfo, err := imageHandler.GetImage(imageId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(imageInfo)
				}
				fmt.Println("Finish GetImage()")
			case 3:
				fmt.Println("Start CreateImage() ...")
				reqInfo := irs.ImageReqInfo{
					IId: irs.IID{
						NameId: config.Openstack.Image.Name,
					},
				}
				image, err := imageHandler.CreateImage(reqInfo)
				if err != nil {
					fmt.Println(err)
				}
				imageId = image.IId
				fmt.Println("Finish CreateImage()")
			case 4:
				fmt.Println("Start DeleteImage() ...")
				if ok, err := imageHandler.DeleteImage(imageId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteImage()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testVPCHandler(config Config) {
	resourceHandler, err := getResourceHandler("vpc")
	if err != nil {
		fmt.Println(err)
	}

	vpcHandler := resourceHandler.(irs.VPCHandler)

	fmt.Println("Test VPCHandler")
	fmt.Println("1. ListVPC()")
	fmt.Println("2. GetVPC()")
	fmt.Println("3. CreateVPC()")
	fmt.Println("4. DeleteVPC()")
	fmt.Println("5. Exit")

	vpcId := irs.IID{NameId: "CB-VNet2", SystemId: "fa517cc1-7d4a-4a6f-a0be-ff77761152b5"}

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVPC() ...")
				if list, err := vpcHandler.ListVPC(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(list)
				}
				fmt.Println("Finish ListVPC()")
			case 2:
				fmt.Println("Start GetVPC() ...")
				if vNetInfo, err := vpcHandler.GetVPC(vpcId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vNetInfo)
				}
				fmt.Println("Finish GetVPC()")
			case 3:
				fmt.Println("Start CreateVPC() ...")
				reqInfo := irs.VPCReqInfo{
					IId: vpcId,
					SubnetInfoList: []irs.SubnetInfo{
						{
							IId: irs.IID{
								NameId: vpcId.NameId + "-subnet-1",
							},
							IPv4_CIDR: "180.0.10.0/24",
						},
						{
							IId: irs.IID{
								NameId: vpcId.NameId + "-subnet-2",
							},
							IPv4_CIDR: "180.0.20.0/24",
						},
					},
				}

				vpcInfo, err := vpcHandler.CreateVPC(reqInfo)
				if err != nil {
					fmt.Println(err)
				}

				vpcId = vpcInfo.IId
				spew.Dump(vpcInfo)
				fmt.Println("Finish CreateVPC()")
			case 4:
				fmt.Println("Start DeleteVPC() ...")
				if ok, err := vpcHandler.DeleteVPC(vpcId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteVPC()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func testKeyPairHandler(config Config) {
	resourceHandler, err := getResourceHandler("keypair")
	if err != nil {
		fmt.Println(err)
	}

	keyPairHandler := resourceHandler.(irs.KeyPairHandler)

	fmt.Println("Test KeyPairHandler")
	fmt.Println("1. ListKey()")
	fmt.Println("2. GetKey()")
	fmt.Println("3. CreateKey()")
	fmt.Println("4. DeleteKey()")
	fmt.Println("5. Exit")

	keypairIId := irs.IID{
		NameId: "CB-Keypair",
	}

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListKey() ...")
				if keyPairList, err := keyPairHandler.ListKey(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(keyPairList)
				}
				fmt.Println("Finish ListKey()")
			case 2:
				fmt.Println("Start GetKey() ...")
				if keyPairInfo, err := keyPairHandler.GetKey(keypairIId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(keyPairInfo)
				}
				fmt.Println("Finish GetKey()")
			case 3:
				fmt.Println("Start CreateKey() ...")
				reqInfo := irs.KeyPairReqInfo{
					IId: keypairIId,
				}
				if keyInfo, err := keyPairHandler.CreateKey(reqInfo); err != nil {
					fmt.Println(err)
				} else {
					keypairIId = keyInfo.IId
					spew.Dump(keyInfo)
				}
				fmt.Println("Finish CreateKey()")
			case 4:
				fmt.Println("Start DeleteKey() ...")
				if ok, err := keyPairHandler.DeleteKey(keypairIId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteKey()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

/*func testPublicIPHanlder(config Config) {
	resourceHandler, err := getResourceHandler("publicip")
	if err != nil {
		fmt.Println(err)
	}

	publicIPHandler := resourceHandler.(irs.PublicIPHandler)

	fmt.Println("Test PublicIPHandler")
	fmt.Println("1. ListPublicIP()")
	fmt.Println("2. GetPublicIP()")
	fmt.Println("3. CreatePublicIP()")
	fmt.Println("4. DeletePublicIP()")
	fmt.Println("5. Exit")

	var publicIPId string

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListPublicIP() ...")
				if publicList, err := publicIPHandler.ListPublicIP(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(publicList)
				}
				fmt.Println("Finish ListPublicIP()")
			case 2:
				fmt.Println("Start GetPublicIP() ...")
				if publicInfo, err := publicIPHandler.GetPublicIP(publicIPId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(publicInfo)
				}
				fmt.Println("Finish GetPublicIP()")
			case 3:
				fmt.Println("Start CreatePublicIP() ...")

				reqInfo := irs.PublicIPReqInfo{}
				if publicIP, err := publicIPHandler.CreatePublicIP(reqInfo); err != nil {
					fmt.Println(err)
				} else {
					publicIPId = publicIP.Name
					spew.Dump(publicIP)
				}
				fmt.Println("Finish CreatePublicIP()")
			case 4:
				fmt.Println("Start DeletePublicIP() ...")
				if ok, err := publicIPHandler.DeletePublicIP(publicIPId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeletePublicIP()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}*/

func testSecurityHandler(config Config) {
	resourceHandler, err := getResourceHandler("security")
	if err != nil {
		fmt.Println(err)
	}

	securityHandler := resourceHandler.(irs.SecurityHandler)

	fmt.Println("Test SecurityHandler")
	fmt.Println("1. ListSecurity()")
	fmt.Println("2. GetSecurity()")
	fmt.Println("3. CreateSecurity()")
	fmt.Println("4. DeleteSecurity()")
	fmt.Println("5. Exit")

	securityGroupIId := irs.IID{
		NameId: "CB-SecGroup",
	}

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListSecurity() ...")
				if securityList, err := securityHandler.ListSecurity(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(securityList)
				}
				fmt.Println("Finish ListSecurity()")
			case 2:
				fmt.Println("Start GetSecurity() ...")
				if secInfo, err := securityHandler.GetSecurity(securityGroupIId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(secInfo)
				}
				fmt.Println("Finish GetSecurity()")
			case 3:
				fmt.Println("Start CreateSecurity() ...")

				reqInfo := irs.SecurityReqInfo{
					IId: irs.IID{
						NameId: securityGroupIId.NameId,
					},
					SecurityRules: &[]irs.SecurityRuleInfo{
						{
							FromPort:   "22",
							ToPort:     "22",
							IPProtocol: "TCP",
							Direction:  "inbound",
						},
						{
							FromPort:   "3306",
							ToPort:     "3306",
							IPProtocol: "TCP",
							Direction:  "outbound",
						},
						{
							IPProtocol: "ICMP",
							Direction:  "outbound",
						},
					},
				}
				if securityInfo, err := securityHandler.CreateSecurity(reqInfo); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(securityInfo)
					securityGroupIId = securityInfo.IId
				}
				fmt.Println("Finish CreateSecurity()")
			case 4:
				fmt.Println("Start DeleteSecurity() ...")
				if ok, err := securityHandler.DeleteSecurity(securityGroupIId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteSecurity()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

/*func testVNetworkHandler(config Config) {
	resourceHandler, err := getResourceHandler("vnetwork")
	if err != nil {
		fmt.Println(err)
	}

	vNetworkHandler := resourceHandler.(irs.VNetworkHandler)

	fmt.Println("Test VNetworkHandler")
	fmt.Println("1. ListVNetwork()")
	fmt.Println("2. GetVNetwork()")
	fmt.Println("3. CreateVNetwork()")
	fmt.Println("4. DeleteVNetwork()")
	fmt.Println("5. Exit")

	vNetWorkName := "CB-VNet"
	var vNetworkId string

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVNetwork() ...")
				if list, err := vNetworkHandler.ListVNetwork(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(list)
				}
				fmt.Println("Finish ListVNetwork()")
			case 2:
				fmt.Println("Start GetVNetwork() ...")
				if vNetInfo, err := vNetworkHandler.GetVNetwork(vNetworkId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vNetInfo)
				}
				fmt.Println("Finish GetVNetwork()")
			case 3:
				fmt.Println("Start CreateVNetwork() ...")

				reqInfo := irs.VNetworkReqInfo{
					Name: vNetWorkName,
				}

				if vNetworkInfo, err := vNetworkHandler.CreateVNetwork(reqInfo); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vNetworkInfo)
					vNetworkId = vNetworkInfo.Id
				}
				fmt.Println("Finish CreateVNetwork()")
			case 4:
				fmt.Println("Start DeleteVNetwork() ...")
				if ok, err := vNetworkHandler.DeleteVNetwork(vNetworkId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteVNetwork()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}*/

/*func testVNicHandler(config Config) {
	resourceHandler, err := getResourceHandler("vnic")
	if err != nil {
		fmt.Println(err)
	}

	vNicHandler := resourceHandler.(irs.VNicHandler)

	fmt.Println("Test VNicHandler")
	fmt.Println("1. ListVNic()")
	fmt.Println("2. GetVNic()")
	fmt.Println("3. CreateVNic()")
	fmt.Println("4. DeleteVNic()")
	fmt.Println("5. Exit")

	vNicName := "CB-VNic"
	var vNicId string

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVNic() ...")
				if List, err := vNicHandler.ListVNic(); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(List)
				}
				fmt.Println("Finish ListVNic()")
			case 2:
				fmt.Println("Start GetVNic() ...")
				if vNicInfo, err := vNicHandler.GetVNic(vNicId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vNicInfo)
				}
				fmt.Println("Finish GetVNic()")
			case 3:
				fmt.Println("Start CreateVNic() ...")

				//todo : port로 맵핑
				reqInfo := irs.VNicReqInfo{
					Name:             vNicName,
					VNetId:           "fe284dbf-e9f4-4add-a03f-9249cc30a2ac",
					SecurityGroupIds: []string{"34585b5e-5ea8-49b5-b38b-0d395689c994", "6d4085c1-e915-487d-9e83-7a5b64f27237"},
					//SubnetId:         "fe284dbf-e9f4-4add-a03f-9249cc30a2ac",
				}

				if vNicInfo, err := vNicHandler.CreateVNic(reqInfo); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vNicInfo)
					vNicId = vNicInfo.Id
				}
				fmt.Println("Finish CreateVNic()")
			case 4:
				fmt.Println("Start DeleteVNic() ...")
				if ok, err := vNicHandler.DeleteVNic(vNicId); !ok {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteVNic()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}*/

/*func testRouterHandler(config Config) {
	resourceHandler, err := getResourceHandler("router")
	if err != nil {
		fmt.Println(err)
	}

	routerHandler := resourceHandler.(osrs.OpenStackRouterHandler)

	fmt.Println("Test RouterHandler")
	fmt.Println("1. ListRouter()")
	fmt.Println("2. GetRouter()")
	fmt.Println("3. CreateRouter()")
	fmt.Println("4. DeleteRouter()")
	fmt.Println("5. AddInterface()")
	fmt.Println("6. DeleteInterface()")
	fmt.Println("7. Exit")

	var routerId string

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListRouter() ...")
				routerHandler.ListRouter()
				fmt.Println("Finish ListRouter()")
			case 2:
				fmt.Println("Start GetRouter() ...")
				routerHandler.GetRouter(routerId)
				fmt.Println("Finish GetRouter()")
			case 3:
				fmt.Println("Start CreateRouter() ...")
				reqInfo := osrs.RouterReqInfo{
					Name:         config.Openstack.Router.Name,
					GateWayId:    config.Openstack.Router.GateWayId,
					AdminStateUp: config.Openstack.Router.AdminStateUp,
				}
				router, err := routerHandler.CreateRouter(reqInfo)
				if err != nil {
					fmt.Println(err)
				}
				routerId = router.Id
				fmt.Println("Finish CreateRouter()")
			case 4:
				fmt.Println("Start DeleteRouter() ...")
				routerHandler.DeleteRouter(routerId)
				fmt.Println("Finish DeleteRouter()")
			case 5:
				fmt.Println("Start AddInterface() ...")
				reqInfo := osrs.InterfaceReqInfo{
					SubnetId: config.Openstack.Subnet.Id,
					RouterId: routerId,
				}
				_, err := routerHandler.AddInterface(reqInfo)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Finish AddInterface()")
			case 6:
				fmt.Println("Start DeleteInterface() ...")
				_, err := routerHandler.DeleteInterface(routerId, config.Openstack.Subnet.Id)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Finish DeleteInterface()")
			case 7:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}*/

func testVMSpecHandler(config Config) {
	resourceHandler, err := getResourceHandler("vmspec")
	if err != nil {
		panic(err)
	}

	vmSpecHandler := resourceHandler.(irs.VMSpecHandler)

	fmt.Println("Test VMSpecHandler")
	fmt.Println("1. ListVMSpec()")
	fmt.Println("2. GetVMSpec()")
	fmt.Println("3. ListOrgVMSpec()")
	fmt.Println("4. GetOrgVMSpec()")
	fmt.Println("5. Exit")

	var vmSpecId string
	vmSpecId = "babo"

Loop:
	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		region := config.Openstack.Region

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				fmt.Println("Start ListVMSpec() ...")
				if list, err := vmSpecHandler.ListVMSpec(region); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(list)
				}
				fmt.Println("Finish ListVMSpec()")
			case 2:
				fmt.Println("Start GetVMSpec() ...")
				if vmSpecInfo, err := vmSpecHandler.GetVMSpec(region, vmSpecId); err != nil {
					fmt.Println(err)
				} else {
					spew.Dump(vmSpecInfo)
				}
				fmt.Println("Finish GetVMSpec()")
			case 3:
				fmt.Println("Start ListOrgVMSpec() ...")
				if listStr, err := vmSpecHandler.ListOrgVMSpec(region); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(listStr)
				}
				fmt.Println("Finish ListOrgVMSpec()")
			case 4:
				fmt.Println("Start GetOrgVMSpec() ...")
				if vmSpecStr, err := vmSpecHandler.GetOrgVMSpec(region, vmSpecId); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(vmSpecStr)
				}
				fmt.Println("Finish GetOrgVMSpec()")
			case 5:
				fmt.Println("Exit")
				break Loop
			}
		}
	}
}

func getResourceHandler(resourceType string) (interface{}, error) {
	var cloudDriver idrv.CloudDriver
	cloudDriver = new(osdrv.OpenStackDriver)

	config := readConfigFile()
	connectionInfo := idrv.ConnectionInfo{
		CredentialInfo: idrv.CredentialInfo{
			IdentityEndpoint: config.Openstack.IdentityEndpoint,
			Username:         config.Openstack.Username,
			Password:         config.Openstack.Password,
			DomainName:       config.Openstack.DomainName,
			ProjectID:        config.Openstack.ProjectID,
		},
		RegionInfo: idrv.RegionInfo{
			Region: config.Openstack.Region,
		},
	}

	cloudConnection, _ := cloudDriver.ConnectCloud(connectionInfo)

	var resourceHandler interface{}
	var err error

	switch resourceType {
	case "image":
		resourceHandler, err = cloudConnection.CreateImageHandler()
	case "keypair":
		resourceHandler, err = cloudConnection.CreateKeyPairHandler()
	//case "publicip":
	//	resourceHandler, err = cloudConnection.CreatePublicIPHandler()
	case "security":
		resourceHandler, err = cloudConnection.CreateSecurityHandler()
	//case "vnetwork":
	//	resourceHandler, err = cloudConnection.CreateVNetworkHandler()
	case "vpc":
		resourceHandler, err = cloudConnection.CreateVPCHandler()
	//case "vnic":
	//	resourceHandler, err = cloudConnection.CreateVNicHandler()
	case "router":
		//osDriver := osdrv.OpenStackDriver{}
		//cloudConn, err := osDriver.ConnectCloud(connectionInfo)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//osCloudConn := cloudConn.(*connect.OpenStackCloudConnection)
		//resourceHandler = osrs.OpenStackRouterHandler{Client: osCloudConn.NetworkClient}
	case "vmspec":
		resourceHandler, err = cloudConnection.CreateVMSpecHandler()
	}

	if err != nil {
		return nil, err
	}
	return resourceHandler, nil
}

func showTestHandlerInfo() {
	fmt.Println("==========================================================")
	fmt.Println("[Test ResourceHandler]")
	fmt.Println("1. ImageHandler")
	fmt.Println("2. KeyPairHandler")
	//fmt.Println("3. PublicIPHandler")
	fmt.Println("4. SecurityHandler")
	fmt.Println("5. VPCHandler")
	//fmt.Println("6. VNicHandler")
	fmt.Println("7. RouterHandler")
	fmt.Println("8. VMSpecHandler")
	fmt.Println("9. Exit")
	fmt.Println("==========================================================")
}

func main() {

	showTestHandlerInfo()      // ResourceHandler 테스트 정보 출력
	config := readConfigFile() // config.yaml 파일 로드

Loop:

	for {
		var commandNum int
		inputCnt, err := fmt.Scan(&commandNum)
		if err != nil {
			fmt.Println(err)
		}

		if inputCnt == 1 {
			switch commandNum {
			case 1:
				testImageHandler(config)
				showTestHandlerInfo()
			case 2:
				testKeyPairHandler(config)
				showTestHandlerInfo()
			case 3:
				//testPublicIPHanlder(config)
				//showTestHandlerInfo()
			case 4:
				testSecurityHandler(config)
				showTestHandlerInfo()
			case 5:
				//testVNetworkHandler(config)
				testVPCHandler(config)
				showTestHandlerInfo()
			case 6:
				//testVNicHandler(config)
				//showTestHandlerInfo()
			case 7:
				//testRouterHandler(config)
				showTestHandlerInfo()
			case 8:
				testVMSpecHandler(config)
				showTestHandlerInfo()
			case 9:
				fmt.Println("Exit Test ResourceHandler Program")
				break Loop
			}
		}
	}
}

type Config struct {
	Openstack struct {
		DomainName       string `yaml:"domain_name"`
		IdentityEndpoint string `yaml:"identity_endpoint"`
		Password         string `yaml:"password"`
		ProjectID        string `yaml:"project_id"`
		Username         string `yaml:"username"`
		Region           string `yaml:"region"`
		VMName           string `yaml:"vm_name"`
		ImageId          string `yaml:"image_id"`
		FlavorId         string `yaml:"flavor_id"`
		NetworkId        string `yaml:"network_id"`
		SecurityGroups   string `yaml:"security_groups"`
		KeypairName      string `yaml:"keypair_name"`

		ServerId string `yaml:"server_id"`

		Image struct {
			Name string `yaml:"name"`
		} `yaml:"image_info"`

		KeyPair struct {
			Name string `yaml:"name"`
		} `yaml:"keypair_info"`

		PublicIP struct {
			Name string `yaml:"name"`
		} `yaml:"public_info"`

		SecurityGroup struct {
			Name string `yaml:"name"`
		} `yaml:"security_group_info"`

		VirtualNetwork struct {
			Name string `yaml:"name"`
		} `yaml:"vnet_info"`

		Subnet struct {
			Id string `yaml:"id"`
		} `yaml:"subnet_info"`

		Router struct {
			Name         string `yaml:"name"`
			GateWayId    string `yaml:"gateway_id"`
			AdminStateUp bool   `yaml:"adminstatup"`
		} `yaml:"router_info"`
	} `yaml:"openstack"`
}

func readConfigFile() Config {
	// Set Environment Value of Project Root Path
	rootPath := os.Getenv("CBSPIDER_ROOT")
	data, err := ioutil.ReadFile(rootPath + "/conf/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
