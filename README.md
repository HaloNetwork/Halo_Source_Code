# HALO Network

 

​	   	Halo Network is an efficient blockchain network system that with improvements and extensions based on Ethereum's open-source protocol and redefine a new protocol called the Halo Network to implement the Infinite Financial System solution.

​	   	Halo Network adopts innovative transaction management with bimodal liquidity, cross-chain swap along with new HPOS consensus mechanism and native oracle machine in the base layer. The improvement of the base layer enables the Halo Network to obtain the basic conditions of accessing the traditional financial instruments and achieve a more efficient and low-cost transaction structure. Halo Network prides itself in its value proposition by seamlessly connecting the traditional financial instruments to the blockchain system and unifying decentralized transactions at all levels.

 

***\*License\****

The ***\*HaloNetwork\**** library (i.e. all code outside of the cmd directory) is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the COPYING.LESSER file.

The ***\*HaloNetwork\**** binaries (i.e. all code inside of the cmd directory) is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the COPYING file.

# Halo Node deployment

## Minimum CPU requirement

4core 
8GB RAM 
300G SSD 
CentOS 7.6

## 1)Setting up the environment

yum -y install wget git gcc-c++

## 2)Download "golang"

wget https://studygolang.com/dl/golang/go1.15.3.linux-amd64.tar.gz

## 3)Unzip "golang" zip file

tar -zxvf go1.15.3.linux-amd64.tar.gz //the zip file directory

## 4)Setup "go" environment variable

vi /etc/profile
export PATH="$PATH:/usr/local/go/bin"

## 5)Save and referesh "profile"

source /etc/profile

## 6)Validate installation successfully

go version
go version go1.15.3 linux/amd64

## 7)Download HALO source code

git clone https://github.com/HaloNetwork/Halo_Source_Code.git

## 8)Navigate to HALO source code directory, then make "gho"

For example:
cd /root/Halo_Source_Code
make gho

## 9)Copy "gho" to "bin"'

cp  build/bin/gho /usr/local/bin/


Preliminary work is ready, now start to synchronize nodes.
The following operations are for reference only
The actual operation subject to user's PC directory and port

## 1)Creating a node directory

mkdir halonode

cd halonode && mkdir data

## 2)initialization  of the node 

gho init --datadir data/ /root/Halo_Source_Code/genesis_halo_prod.json

## 3)Start the node after initialization is completed

gho --datadir /root/halonode/data --port "20202" --http --http.addr "0.0.0.0" --http.port "8845" --http.corsdomain "*" --http.vhosts "*" --ws --ws.addr "0.0.0.0" --ws.port "8846" --ws.origins "*" 