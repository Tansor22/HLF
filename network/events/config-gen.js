/**
 * reads the nw-config yaml
 * Replaces the PK/Cert for the peers
 * CODE Replicated in events.js
 */
const fs = require('fs');
const yaml = require('node-yaml')

const CRYPTOGEN_PEER="../crypto/crypto-config/peerOrganizations"

var obj = yaml.readSync('./nw-config.template.yaml')
createYAMLCryptogen(obj)
console.log('done')

// console.log("obj=", obj)



// console.log(genCertPathCryptogen('astu-service'))
// console.log(genPkCryptogen("astu-service"))

// Generates the YAML based on the 
function createYAMLCryptogen(obj){
    let astuCert = genCertPathCryptogen('astu')
    let astuPk = genPkCryptogen("astu")
    obj.organizations.Astu.signedCert.path=astuCert
    obj.organizations.Astu.adminPrivateKey.path=astuPk
    let astu-serviceCert = genCertPathCryptogen('astu-service')
    let astu-servicePk = genPkCryptogen("astu-service")
    obj.organizations.Astu-Service.signedCert.path=astu-serviceCert
    obj.organizations.Astu-Service.adminPrivateKey.path=astu-servicePk

    console.log(astuCert)
    yaml.writeSync("nw-config.yaml", obj)
}

function genCertPathCryptogen(org){ 
    //astu-service.com/users/Admin@astu-service.com/msp/signcerts/Admin@astu-service.com-cert.pem"
    var certPath=CRYPTOGEN_PEER+"/"+org+".com/users/Admin@"+org+".com/msp/signcerts/Admin@"+org+".com-cert.pem"
    return certPath
}

// looks for the PK files in the org folder
function    genPkCryptogen(org){
    // ../crypto/crypto-config/peerOrganizations/astu-service.com/users/Admin@astu-service.com/msp/keystore/05beac9849f610ad5cc8997e5f45343ca918de78398988def3f288b60d8ee27c_sk
    var pkFolder=CRYPTOGEN_PEER+"/"+org+".com/users/Admin@"+org+".com/msp/keystore"
    fs.readdirSync(pkFolder).forEach(file => {
        // console.log(file);
        // return the first file
        pkfile = file
        return
    })

    return pkFolder+"/"+pkfile
}