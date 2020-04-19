// Copyright 2020 Cryptlex, LLC. All rights reserved.

package lexactivator

/*
#cgo LDFLAGS: -L./libs -lLexActivator -Wl,-rpath,./libs -Wl,-rpath,./
#include "lexactivator/LexActivator.h"
#include <stdlib.h>
void licenseCallbackCgoGateway(int status);
void releaseUpdateCallbackCgoGateway(int status);
*/
import "C"
import (
	"unsafe"
)

type CallbackType func(int)

var licenseCallbackFuncion CallbackType

var releaseUpdateCallbackFuncion CallbackType

const (
	LA_USER      uint = 0
	LA_SYSTEM    uint = 1
	LA_IN_MEMORY uint = 2
)

//export licenseCallbackWrapper
func licenseCallbackWrapper(status int) {
	if licenseCallbackFuncion != nil {
		licenseCallbackFuncion(status)
	}
}

//export releaseUpdateCallbackWrapper
func releaseUpdateCallbackWrapper(status int) {
	if releaseUpdateCallbackFuncion != nil {
		releaseUpdateCallbackFuncion(status)
	}
}


/*
    FUNCTION: SetProductFile()

    PURPOSE: Sets the absolute path of the Product.dat file.

    This function must be called on every start of your program
    before any other functions are called.

    PARAMETERS:
    * filePath - absolute path of the product file (Product.dat)

    RETURN CODES: LA_OK, LA_E_FILE_PATH, LA_E_PRODUCT_FILE

    NOTE: If this function fails to set the path of product file, none of the
    other functions will work.
*/
func SetProductFile(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.SetProductFile(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: SetProductData()

    PURPOSE: Embeds the Product.dat file in the application.

    It can be used instead of SetProductFile() in case you want
    to embed the Product.dat file in your application.

    This function must be called on every start of your program
    before any other functions are called.

    PARAMETERS:
    * productData - content of the Product.dat file

    RETURN CODES: LA_OK, LA_E_PRODUCT_DATA

    NOTE: If this function fails to set the product data, none of the
    other functions will work.
*/
func SetProductData(productData string) int {
	cProductData := GoToCString(productData)
	status := C.SetProductData(cProductData)
	FreeCString(cProductData)
	return int(status)
}

/*
    FUNCTION: SetProductId()

    PURPOSE: Sets the product id of your application.

    This function must be called on every start of your program before
    any other functions are called, with the exception of SetProductFile()
    or SetProductData() function.

    PARAMETERS:
    * productId - the unique product id of your application as mentioned
      on the product page in the dashboard.

    * flags - depending upon whether your application requires admin/root
      permissions to run or not, this parameter can have one of the following
      values: LA_SYSTEM, LA_USER, LA_IN_MEMORY

    RETURN CODES: LA_OK, LA_E_WMIC, LA_E_PRODUCT_FILE, LA_E_PRODUCT_DATA, LA_E_PRODUCT_ID,
    LA_E_SYSTEM_PERMISSION

    NOTE: If this function fails to set the product id, none of the other
    functions will work.
*/
func SetProductId(productId string, flags uint) int {
	cProductId := GoToCString(productId)
	cFlags := (C.uint)(flags)
	status := C.SetProductId(cProductId, cFlags)
	FreeCString(cProductId)
	return int(status)
}

/*
    FUNCTION: SetLicenseKey()

    PURPOSE: Sets the license key required to activate the license.

    PARAMETERS:
    * licenseKey - a valid license key.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseKey(licenseKey string) int {
	cLicenseKey := GoToCString(licenseKey)
	status := C.SetLicenseKey(cLicenseKey)
	FreeCString(cLicenseKey)
	return int(status)
}

/*
    FUNCTION: SetLicenseUserCredential()

    PURPOSE: Sets the license user email and password for authentication.

    This function must be called before ActivateLicense() or IsLicenseGenuine()
    function if 'requireAuthentication' property of the license is set to true.

    PARAMETERS:
    * email - user email address.
    * password - user password.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseUserCredential(email string, password string) int {
	cEmail := GoToCString(email)
	cPassword := GoToCString(password)
	status := C.SetLicenseUserCredential(cEmail, cPassword)
	FreeCString(cEmail)
	FreeCString(cPassword)
	return int(status)
}

/*
    FUNCTION: SetLicenseCallback()

    PURPOSE: Sets server sync callback function.

    Whenever the server sync occurs in a separate thread, and server returns the response,
    license callback function gets invoked with the following status codes:
    LA_OK, LA_EXPIRED, LA_SUSPENDED,
    LA_E_REVOKED, LA_E_ACTIVATION_NOT_FOUND, LA_E_MACHINE_FINGERPRINT
    LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_INET, LA_E_SERVER,
    LA_E_RATE_LIMIT, LA_E_IP

    PARAMETERS:
    * callback - name of the callback function

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY
*/
func SetLicenseCallback(callbackFunction CallbackType) int {
	status := C.SetLicenseCallback((C.CallbackType)(unsafe.Pointer(C.licenseCallbackCgoGateway)))
	licenseCallbackFuncion = callbackFunction
	return int(status)
}

/*
    FUNCTION: SetActivationMetadata()

    PURPOSE: Sets the activation metadata.

    The  metadata appears along with the activation details of the license
    in dashboard.

    PARAMETERS:
    * key - string of maximum length 256 characters with utf-8 encoding.
    * value - string of maximum length 256 characters with utf-8 encoding.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_METADATA_KEY_LENGTH,
    LA_E_METADATA_VALUE_LENGTH, LA_E_ACTIVATION_METADATA_LIMIT
*/
func SetActivationMetadata(key string, value string) int {
	cKey := GoToCString(key)
	cValue := GoToCString(value)
	status := C.SetActivationMetadata(cKey, cValue)
	FreeCString(cKey)
	FreeCString(cValue)
	return int(status)
}

/*
    FUNCTION: SetTrialActivationMetadata()

    PURPOSE: Sets the trial activation metadata.

    The  metadata appears along with the trial activation details of the product
    in dashboard.

    PARAMETERS:
    * key - string of maximum length 256 characters with utf-8 encoding.
    * value - string of maximum length 256 characters with utf-8 encoding.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_LENGTH,
    LA_E_METADATA_VALUE_LENGTH, LA_E_TRIAL_ACTIVATION_METADATA_LIMIT
*/
func SetTrialActivationMetadata(key string, value string) int {
	cKey := GoToCString(key)
	cValue := GoToCString(value)
	status := C.SetTrialActivationMetadata(cKey, cValue)
	FreeCString(cKey)
	FreeCString(cValue)
	return int(status)
}

/*
    FUNCTION: SetAppVersion()

    PURPOSE: Sets the current app version of your application.

    The app version appears along with the activation details in dashboard. It
    is also used to generate app analytics.

    PARAMETERS:
    * appVersion - string of maximum length 256 characters with utf-8 encoding.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_APP_VERSION_LENGTH
*/
func SetAppVersion(appVersion string) int {
	cAppVersion := GoToCString(appVersion)
	status := C.SetAppVersion(cAppVersion)
	FreeCString(cAppVersion)
	return int(status)
}

/*
    FUNCTION: SetOfflineActivationRequestMeterAttributeUses()

    PURPOSE: Sets the meter attribute uses for the offline activation request.

    This function should only be called before GenerateOfflineActivationRequest()
    function to set the meter attributes in case of offline activation.

    PARAMETERS:
    * name - name of the meter attribute
    * uses - the uses value

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY

*/
func SetOfflineActivationRequestMeterAttributeUses(name string, uses uint) int {
	cName := GoToCString(name)
	cUses := (C.uint)(uses)
	status := C.SetOfflineActivationRequestMeterAttributeUses(cName, cUses)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: SetNetworkProxy()

    PURPOSE: Sets the network proxy to be used when contacting Cryptlex servers.

    The proxy format should be: [protocol://][username:password@]machine[:port]

    Following are some examples of the valid proxy strings:
        - http://127.0.0.1:8000/
        - http://user:pass@127.0.0.1:8000/
        - socks5://127.0.0.1:8000/

    PARAMETERS:
    * proxy - proxy string having correct proxy format

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_NET_PROXY

    NOTE: Proxy settings of the computer are automatically detected. So, in most of the
    cases you don't need to care whether your user is behind a proxy server or not.
*/
func SetNetworkProxy(proxy string) int {
	cProxy := GoToCString(proxy)
	status := C.SetNetworkProxy(cProxy)
	FreeCString(cProxy)
	return int(status)
}

/*
    FUNCTION: SetCryptlexHost()

    PURPOSE: In case you are running Cryptlex on-premise, you can set the
    host for your on-premise server.

    PARAMETERS:
    * host - the address of the Cryptlex on-premise server

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_HOST_URL
*/
func SetCryptlexHost(host string) int {
	cHost := GoToCString(host)
	status := C.SetCryptlexHost(cHost)
	FreeCString(cHost)
	return int(status)
}

/*
    FUNCTION: GetProductMetadata()

    PURPOSE: Gets the product metadata as set in the dashboard.

    This is available for trial as well as license activations.

    PARAMETERS:
    * key - metadata key to retrieve the value
    * value - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetProductMetadata(key string, value *string) int {
	cKey := GoToCString(key)
	var cValue = GetCArray()
	status := C.GetProductMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = CtoGoString(&cValue[0])
	FreeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetLicenseMetadata()

    PURPOSE: Gets the license metadata as set in the dashboard.

    PARAMETERS:
    * key - metadata key to retrieve the value
    * value - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetLicenseMetadata(key string, value *string) int {
	cKey := GoToCString(key)
	var cValue = GetCArray()
	status := C.GetLicenseMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = CtoGoString(&cValue[0])
	FreeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetLicenseMeterAttribute()

    PURPOSE: Gets the license meter attribute allowed uses and total uses.

    PARAMETERS:
    * name - name of the meter attribute
    * allowedUses - pointer to the integer that receives the value
    * totalUses - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetLicenseMeterAttribute(name string, allowedUses *uint, totalUses *uint) int {
	cName := GoToCString(name)
	var cAllowedUses C.uint
	var cTotalUses C.uint
	status := C.GetLicenseMeterAttribute(cName, &cAllowedUses, &cTotalUses)
	*allowedUses = uint(cAllowedUses)
	*totalUses = uint(cTotalUses)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: GetLicenseKey()

    PURPOSE: Gets the license key used for activation.

    PARAMETERS:
    * licenseKey - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_BUFFER_SIZE
*/
func GetLicenseKey(licenseKey *string) int {
	var cLicenseKey = GetCArray()
	status := C.GetLicenseKey(&cLicenseKey[0], MaxCArrayLength)
	*licenseKey = CtoGoString(&cLicenseKey[0])
	return int(status)
}

/*
    FUNCTION: GetLicenseExpiryDate()

    PURPOSE: Gets the license expiry date timestamp.

    PARAMETERS:
    * expiryDate - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetLicenseExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetLicenseExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

/*
    FUNCTION: GetLicenseUserEmail()

    PURPOSE: Gets the email associated with license user.

    PARAMETERS:
    * email - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetLicenseUserEmail(email *string) int {
	var cEmail = GetCArray()
	status := C.GetLicenseUserEmail(&cEmail[0], MaxCArrayLength)
	*email = CtoGoString(&cEmail[0])
	return int(status)
}

/*
    FUNCTION: GetLicenseUserName()

    PURPOSE: Gets the name associated with the license user.

    PARAMETERS:
    * name - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetLicenseUserName(name *string) int {
	var cName = GetCArray()
	status := C.GetLicenseUserName(&cName[0], MaxCArrayLength)
	*name = CtoGoString(&cName[0])
	return int(status)
}

/*
    FUNCTION: GetLicenseUserCompany()

    PURPOSE: Gets the company associated with the license user.

    PARAMETERS:
    * company - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetLicenseUserCompany(company *string) int {
	var cCompany = GetCArray()
	status := C.GetLicenseUserCompany(&cCompany[0], MaxCArrayLength)
	*company = CtoGoString(&cCompany[0])
	return int(status)
}

/*
    FUNCTION: GetLicenseUserMetadata()

    PURPOSE: Gets the metadata associated with the license user.

    PARAMETERS:
    * key - metadata key to retrieve the value
    * value - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetLicenseUserMetadata(key string, value *string) int {
	cKey := GoToCString(key)
	var cValue = GetCArray()
	status := C.GetLicenseUserMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = CtoGoString(&cValue[0])
	FreeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetLicenseType()

    PURPOSE: Gets the license type (node-locked or hosted-floating).

    PARAMETERS:
    * licenseType - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetLicenseType(licenseType *string) int {
	var cLicenseType = GetCArray()
	status := C.GetLicenseType(&cLicenseType[0], MaxCArrayLength)
	*licenseType = CtoGoString(&cLicenseType[0])
	return int(status)
}

/*
    FUNCTION: GetActivationMetadata()

    PURPOSE: Gets the activation metadata.

    PARAMETERS:
    * key - metadata key to retrieve the value
    * value - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetActivationMetadata(key string, value *string) int {
	cKey := GoToCString(key)
	var cValue = GetCArray()
	status := C.GetActivationMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = CtoGoString(&cValue[0])
	FreeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetActivationMeterAttributeUses()

    PURPOSE: Gets the meter attribute uses consumed by the activation.

    PARAMETERS:
    * name - name of the meter attribute
    * allowedUses - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND
*/
func GetActivationMeterAttributeUses(name string, uses *uint) int {
	cName := GoToCString(name)
	var cUses C.uint
	status := C.GetActivationMeterAttributeUses(cName, &cUses)
	*uses = uint(cUses)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: GetServerSyncGracePeriodExpiryDate()

    PURPOSE: Gets the server sync grace period expiry date timestamp.

    PARAMETERS:
    * expiryDate - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetServerSyncGracePeriodExpiryDate(expiryDate *uint) int {
	var cExpiryDate C.uint
	status := C.GetServerSyncGracePeriodExpiryDate(&cExpiryDate)
	*expiryDate = uint(cExpiryDate)
	return int(status)
}

/*
    FUNCTION: GetTrialActivationMetadata()

    PURPOSE: Gets the trial activation metadata.

    PARAMETERS:
    * key - metadata key to retrieve the value
    * value - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_METADATA_KEY_NOT_FOUND, LA_E_BUFFER_SIZE
*/
func GetTrialActivationMetadata(key string, value *string) int {
	cKey := GoToCString(key)
	var cValue = GetCArray()
	status := C.GetTrialActivationMetadata(cKey, &cValue[0], MaxCArrayLength)
	*value = CtoGoString(&cValue[0])
	FreeCString(cKey)
	return int(status)
}

/*
    FUNCTION: GetTrialExpiryDate()

    PURPOSE: Gets the trial expiry date timestamp.

    PARAMETERS:
    * trialExpiryDate - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GetTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

/*
    FUNCTION: GetTrialId()

    PURPOSE: Gets the trial activation id. Used in case of trial extension.

    PARAMETERS:
    * trialId - pointer to a string that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME, LA_E_TIME_MODIFIED,
    LA_E_BUFFER_SIZE
*/
func GetTrialId(trialId *string) int {
	var cTrialId = GetCArray()
	status := C.GetTrialId(&cTrialId[0], MaxCArrayLength)
	*trialId = CtoGoString(&cTrialId[0])
	return int(status)
}

/*
    FUNCTION: GetLocalTrialExpiryDate()

    PURPOSE: Gets the trial expiry date timestamp.

    PARAMETERS:
    * trialExpiryDate - pointer to the integer that receives the value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED
*/
func GetLocalTrialExpiryDate(trialExpiryDate *uint) int {
	var cTrialExpiryDate C.uint
	status := C.GetLocalTrialExpiryDate(&cTrialExpiryDate)
	*trialExpiryDate = uint(cTrialExpiryDate)
	return int(status)
}

/*
    FUNCTION: CheckForReleaseUpdate()

    PURPOSE: Checks whether a new release is available for the product.

    This function should only be used if you manage your releases through
    Cryptlex release management API.

    PARAMETERS:
    * platform - release platform e.g. windows, macos, linux
    * version - current release version
    * channel - release channel e.g. stable
    * releaseUpdateCallback - name of the callback function.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_RELEASE_VERSION_FORMAT
*/
func CheckForReleaseUpdate(platform string, version string, channel string, callbackFunction CallbackType) int {
	cPlatform := GoToCString(platform)
	cVersion := GoToCString(version)
	cChannel := GoToCString(channel)
	status := C.CheckForReleaseUpdate(cPlatform, cVersion, cChannel, (C.CallbackType)(unsafe.Pointer(C.releaseUpdateCallbackCgoGateway)))
	releaseUpdateCallbackFuncion = callbackFunction
	FreeCString(cPlatform)
	FreeCString(cVersion)
	FreeCString(cChannel)
	return int(status)
}

/*
    FUNCTION: ActivateLicense()

    PURPOSE: Activates the license by contacting the Cryptlex servers. It
    validates the key and returns with encrypted and digitally signed token
    which it stores and uses to activate your application.

    This function should be executed at the time of registration, ideally on
    a button click.

    RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_E_REVOKED, LA_FAIL, LA_E_PRODUCT_ID,
    LA_E_INET, LA_E_VM, LA_E_TIME, LA_E_ACTIVATION_LIMIT, LA_E_SERVER, LA_E_CLIENT,
    LA_E_AUTHENTICATION_FAILED, LA_E_LICENSE_TYPE, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY
*/
func ActivateLicense() int {
	status := C.ActivateLicense()
	return int(status)
}

/*
    FUNCTION: ActivateLicenseOffline()

    PURPOSE: Activates your licenses using the offline activation response file.

    PARAMETERS:
    * filePath - path of the offline activation response file.

    RETURN CODES: LA_OK, LA_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_OFFLINE_RESPONSE_FILE
    LA_E_VM, LA_E_TIME, LA_E_FILE_PATH, LA_E_OFFLINE_RESPONSE_FILE_EXPIRED
*/
func ActivateLicenseOffline(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.ActivateLicenseOffline(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: GenerateOfflineActivationRequest()

    PURPOSE: Generates the offline activation request needed for generating
    offline activation response in the dashboard.

    PARAMETERS:
    * filePath - path of the file for the offline request.

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_FILE_PERMISSION
*/
func GenerateOfflineActivationRequest(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.GenerateOfflineActivationRequest(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: DeactivateLicense()

    PURPOSE: Deactivates the license activation and frees up the corresponding activation
    slot by contacting the Cryptlex servers.

    This function should be executed at the time of de-registration, ideally on
    a button click.

    RETURN CODES: LA_OK, LA_E_DEACTIVATION_LIMIT, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME
    LA_E_LICENSE_KEY, LA_E_INET, LA_E_SERVER, LA_E_RATE_LIMIT, LA_E_TIME_MODIFIED
*/
func DeactivateLicense() int {
	status := C.DeactivateLicense()
	return int(status)
}

/*
    FUNCTION: GenerateOfflineDeactivationRequest()

    PURPOSE: Generates the offline deactivation request needed for deactivation of
    the license in the dashboard and deactivates the license locally.

    A valid offline deactivation file confirms that the license has been successfully
    deactivated on the user's machine.

    PARAMETERS:
    * filePath - path of the file for the offline request.

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_FILE_PERMISSION,
    LA_E_TIME, LA_E_TIME_MODIFIED
*/
func GenerateOfflineDeactivationRequest(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.GenerateOfflineDeactivationRequest(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: IsLicenseGenuine()

    PURPOSE: It verifies whether your app is genuinely activated or not. The verification is
    done locally by verifying the cryptographic digital signature fetched at the time of
    activation.

    After verifying locally, it schedules a server check in a separate thread. After the
    first server sync it periodically does further syncs at a frequency set for the license.

    In case server sync fails due to network error, and it continues to fail for fixed
    number of days (grace period), the function returns LA_GRACE_PERIOD_OVER instead of LA_OK.

    This function must be called on every start of your program to verify the activation
    of your app.

    RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_GRACE_PERIOD_OVER, LA_FAIL,
    LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME, LA_E_TIME_MODIFIED

    NOTE: If application was activated offline using ActivateLicenseOffline() function, you
    may want to set grace period to 0 to ignore grace period.
*/
func IsLicenseGenuine() int {
	status := C.IsLicenseGenuine()
	return int(status)
}

/*
    FUNCTION: IsLicenseValid()

    PURPOSE: It verifies whether your app is genuinely activated or not. The verification is
    done locally by verifying the cryptographic digital signature fetched at the time of
    activation.

    This is just an auxiliary function which you may use in some specific cases, when you
    want to skip the server sync.

    RETURN CODES: LA_OK, LA_EXPIRED, LA_SUSPENDED, LA_GRACE_PERIOD_OVER, LA_FAIL,
    LA_E_PRODUCT_ID, LA_E_LICENSE_KEY, LA_E_TIME, LA_E_TIME_MODIFIED

    NOTE: You may want to set grace period to 0 to ignore grace period.
*/
func IsLicenseValid() int {
	status := C.IsLicenseValid()
	return int(status)
}

/*
    FUNCTION: ActivateTrial()

    PURPOSE: Starts the verified trial in your application by contacting the
    Cryptlex servers.

    This function should be executed when your application starts first time on
    the user's computer, ideally on a button click.

    RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_INET,
    LA_E_VM, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT
*/
func ActivateTrial() int {
	status := C.ActivateTrial()
	return int(status)
}

/*
    FUNCTION: ActivateTrialOffline()

    PURPOSE: Activates your trial using the offline activation response file.

    PARAMETERS:
    * filePath - path of the offline activation response file.

    RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_OFFLINE_RESPONSE_FILE
    LA_E_VM, LA_E_TIME, LA_E_FILE_PATH, LA_E_OFFLINE_RESPONSE_FILE_EXPIRED
*/
func ActivateTrialOffline(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.ActivateTrialOffline(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: GenerateOfflineTrialActivationRequest()

    PURPOSE: Generates the offline trial activation request needed for generating
    offline trial activation response in the dashboard.

    PARAMETERS:
    * filePath - path of the file for the offline request.

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_FILE_PERMISSION
*/
func GenerateOfflineTrialActivationRequest(filePath string) int {
	cFilePath := GoToCString(filePath)
	status := C.GenerateOfflineTrialActivationRequest(cFilePath)
	FreeCString(cFilePath)
	return int(status)
}

/*
    FUNCTION: IsTrialGenuine()

    PURPOSE: It verifies whether trial has started and is genuine or not. The
    verification is done locally by verifying the cryptographic digital signature
    fetched at the time of trial activation.

    This function must be called on every start of your program during the trial period.

    RETURN CODES: LA_OK, LA_TRIAL_EXPIRED, LA_FAIL, LA_E_TIME, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

*/
func IsTrialGenuine() int {
	status := C.IsTrialGenuine()
	return int(status)
}

/*
    FUNCTION: ActivateLocalTrial()

    PURPOSE: Starts the local(unverified) trial.

    This function should be executed when your application starts first time on
    the user's computer.

    PARAMETERS:
    * trialLength - trial length in days

    RETURN CODES: LA_OK, LA_LOCAL_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

    NOTE: The function is only meant for local(unverified) trials.
*/
func ActivateLocalTrial(trialLength uint) int {
	cTrialLength := (C.uint)(trialLength)
	status := C.ActivateLocalTrial(cTrialLength)
	return int(status)
}

/*
    FUNCTION: IsLocalTrialGenuine()

    PURPOSE: It verifies whether trial has started and is genuine or not. The
    verification is done locally.

    This function must be called on every start of your program during the trial period.

    RETURN CODES: LA_OK, LA_LOCAL_TRIAL_EXPIRED, LA_FAIL, LA_E_PRODUCT_ID,
    LA_E_TIME_MODIFIED

    NOTE: The function is only meant for local(unverified) trials.
*/
func IsLocalTrialGenuine() int {
	status := C.IsLocalTrialGenuine()
	return int(status)
}

/*
    FUNCTION: ExtendLocalTrial()

    PURPOSE: Extends the local trial.

    PARAMETERS:
    * trialExtensionLength - number of days to extend the trial

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_TIME_MODIFIED

    NOTE: The function is only meant for local(unverified) trials.
*/
func ExtendLocalTrial(trialExtensionLength uint) int {
	cTrialExtensionLength := (C.uint)(trialExtensionLength)
	status := C.ExtendLocalTrial(cTrialExtensionLength)
	return int(status)
}

/*
    FUNCTION: IncrementActivationMeterAttributeUses()

    PURPOSE: Increments the meter attribute uses of the activation.

    PARAMETERS:
    * name - name of the meter attribute
    * increment - the increment value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
    LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_METER_ATTRIBUTE_USES_LIMIT_REACHED,
    LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY

*/
func IncrementActivationMeterAttributeUses(name string, increment uint) int {
	cName := GoToCString(name)
	cIncrement := (C.uint)(increment)
	status := C.IncrementActivationMeterAttributeUses(cName, cIncrement)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: DecrementActivationMeterAttributeUses()

    PURPOSE: Decrements the meter attribute uses of the activation.

    PARAMETERS:
    * name - name of the meter attribute
    * decrement - the decrement value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
    LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY,
    LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_ACTIVATION_NOT_FOUND

    NOTE: If the decrement is more than the current uses, it resets the uses to 0.
*/
func DecrementActivationMeterAttributeUses(name string, decrement uint) int {
	cName := GoToCString(name)
	cDecrement := (C.uint)(decrement)
	status := C.DecrementActivationMeterAttributeUses(cName, cDecrement)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: ResetActivationMeterAttributeUses()

    PURPOSE: Resets the meter attribute uses consumed by the activation.

    PARAMETERS:
    * name - name of the meter attribute
    * decrement - the decrement value

    RETURN CODES: LA_OK, LA_FAIL, LA_E_PRODUCT_ID, LA_E_METER_ATTRIBUTE_NOT_FOUND,
    LA_E_INET, LA_E_TIME, LA_E_SERVER, LA_E_CLIENT, LA_E_RATE_LIMIT, LA_E_LICENSE_KEY,
    LA_E_AUTHENTICATION_FAILED, LA_E_COUNTRY, LA_E_IP, LA_E_ACTIVATION_NOT_FOUND
*/
func ResetActivationMeterAttributeUses(name string) int {
	cName := GoToCString(name)
	status := C.ResetActivationMeterAttributeUses(cName)
	FreeCString(cName)
	return int(status)
}

/*
    FUNCTION: Reset()

    PURPOSE: Resets the activation and trial data stored in the machine.

    This function is meant for developer testing only.

    RETURN CODES: LA_OK, LA_E_PRODUCT_ID

    NOTE: The function does not reset local(unverified) trial data.
*/
func Reset() int {
	status := C.Reset()
	return int(status)
}
