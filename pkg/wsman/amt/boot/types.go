/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type SettingData struct {
	base message.Base
}

type Capabilities struct {
	base message.Base
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                     xml.Name `xml:"Body"`
		BootSettingDataGetResponse  BootSettingDataResponse
		BootCapabilitiesGetResponse BootCapabilitiesResponse
		EnumerateResponse           common.EnumerateResponse
		PullResponse                PullResponse
	}
	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		BootSettingDataItems  []BootSettingDataResponse  `xml:"Items>AMT_BootSettingData"`
		BootCapabilitiesItems []BootCapabilitiesResponse `xml:"Items>AMT_BootCapabilities"`
	}
	BootSettingDataResponse struct {
		XMLName                  xml.Name          `xml:"AMT_BootSettingData"`
		InstanceID               string            `xml:"InstanceID,omitempty"`               // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>.  <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		ElementName              string            `xml:"ElementName,omitempty"`              // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		OwningEntity             string            `xml:"OwningEntity,omitempty"`             // OwningEntity identifies the vendor or organization that defines the contained boot settings.
		UseSOL                   bool              `xml:"UseSOL,omitempty"`                   // When True, Serial over LAN is used on the next boot cycle.
		UseSafeMode              bool              `xml:"UseSafeMode,omitempty"`              // When a Hard-drive boot source is chosen (using CIM_BootConfigSetting) and this property is set to True, the Intel® AMT firmware will boot in safe mode.
		ReflashBIOS              bool              `xml:"ReflashBIOS,omitempty"`              // When True, the Intel® AMT firmware reflashes the BIOS on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
		BIOSSetup                bool              `xml:"BIOSSetup,omitempty"`                // When True, the Intel® AMT firmware enters the CMOS Setup screen on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
		BIOSPause                bool              `xml:"BIOSPause,omitempty"`                // When True, the BIOS pauses for user input on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
		LockPowerButton          bool              `xml:"LockPowerButton,omitempty"`          // When True, the Intel® AMT firmware disables the power button operation for the system, normally until the next boot cycle.
		LockResetButton          bool              `xml:"LockResetButton,omitempty"`          // When True, the Intel® AMT firmware disables the reset button operation for the system, normally until the next boot cycle.
		LockKeyboard             bool              `xml:"LockKeyboard,omitempty"`             // When True, the Intel® AMT firmware disallows keyboard activity during its boot process.
		LockSleepButton          bool              `xml:"LockSleepButton,omitempty"`          // When True, the Intel® AMT firmware disables the sleep button operation for the system, normally until the next boot cycle.
		UserPasswordBypass       bool              `xml:"UserPasswordBypass,omitempty"`       // When True, the Intel® AMT firmware boots the system and bypasses any user or boot password that might be set in the system.
		ForcedProgressEvents     bool              `xml:"ForcedProgressEvents,omitempty"`     // When True, the Intel® AMT firmware transmits all progress PET events to the alert-sending device.
		FirmwareVerbosity        FirmwareVerbosity `xml:"FirmwareVerbosity,omitempty"`        // When set to a non-zero value, controls the amount of information the managed system writes to its local display.
		ConfigurationDataReset   bool              `xml:"ConfigurationDataReset,omitempty"`   // When True, the Intel® AMT firmware resets its non-volatile configuration data to the managed system's Setup defaults prior to booting the system.
		IDERBootDevice           IDERBootDevice    `xml:"IDERBootDevice,omitempty"`           // Specifies the device to use when UseIder is set. 0 - Floppy Boot, 1- CD Boot.
		UseIDER                  bool              `xml:"UseIDER,omitempty"`                  // When True, IDER is used on the next boot cycle.
		EnforceSecureBoot        bool              `xml:"EnforceSecureBoot,omitempty"`        // When True, Secure boot over IDER is enforced on the next boot cycle, if IDER boot is used. This field is also used in One-Click Recovery.
		BootMediaIndex           int               `xml:"BootMediaIndex,omitempty"`           // This property identifies the boot-media index for the managed client (when a boot source is set using the CIM_BootConfigSetting.ChangeBootOrder method). For Hard-Drive or CD/DVD boot - when the parameter value is 0, the default boot-media is booted. When the parameter value is 1, the primary boot-media is booted; when the value is 2, the secondary boot-media is booted; and so on. For PXE or diagnostics boot this property must be 0.
		SecureErase              bool              `xml:"SecureErase,omitempty"`              // When True, the BIOS performs secure erase operation. Note: Customers are recommended to use Secure Remote Platform Erase which is newer and more advanced than this function.
		RSEPassword              string            `xml:"RSEPassword,omitempty"`              // SSD password for Remote Secure Erase operation. This is a write-only field, an empty string is returned when instance is read. When writing, an empty string or lack of field will be ignored. The password length is limited to 32 ASCII characters. Note: Customers are recommended to use Secure Remote Platform Erase which is newer and more advanced than Remote Secure Erase.
		OptionsCleared           bool              `xml:"OptionsCleared,omitempty"`           // Indicates whether the boot options have been cleared by BIOS or not. This property is read only.
		WinREBootEnabled         bool              `xml:"WinREBootEnabled,omitempty"`         // Indicates whether ForceWinREBoot is enabled in BIOS. This property is read only.
		UEFILocalPBABootEnabled  bool              `xml:"UEFILocalPBABootEnabled,omitempty"`  // Indicates whether ForceUEFILocalPBABoot is enabled in BIOS. This property is read only.
		UEFIHTTPSBootEnabled     bool              `xml:"UEFIHTTPSBootEnabled,omitempty"`     // Indicates whether ForceUEFIHTTPSBoot is enabled in BIOS. This property is read only.
		SecureBootControlEnabled bool              `xml:"SecureBootControlEnabled,omitempty"` // Determines whether Intel AMT is privileged by BIOS to disable secure boot for an AMT triggered boot option. If not, BIOSSecureBoot must be set to TRUE. This property is read only.
		BootguardStatus          bool              `xml:"BootguardStatus,omitempty"`          // Enables the console to discover the security level of the BIOS boot flow. This property is read only.
		BIOSLastStatus           []int             `xml:"BIOSLastStatus,omitempty"`           // Last boot status reported by BIOS. The first 16-bit word contains the general BIOS status (0 - Success, 1 - In Progress, 2 - Not Updated, 0xFFFF - Failed). The second word contains the detailed error status (0 - Success/In Progress, 1 - General Drive Failure, 2 - Drive Password/Authentication Failure, 3 - Feature is not supported). This property is read-only.
		UEFIBootParametersArray  []int             `xml:"UEFIBootParametersArray,omitempty"`  // TLV parameters array encoded with base64 for configuring boot parameters for One-Click Recovery and Secure Remote Platform Erase.
		UEFIBootNumberOfParams   []int             `xml:"UEFIBootNumberOfParams,omitempty"`   // Number of parameters in UefiBootParametersArray
		RPEEnabled               bool              `xml:"RPEEnabled,omitempty"`               // Indicates whether Secure Remote Platform Erase is enabled by the BIOS. Note: This command needs to execute over TLS.
		PlatformErase            bool              `xml:"PlatformErase,omitempty"`            // When set to True, sets the boot option to trigger Secure Remote Platform Erase in the next boot.  Note: This command needs to execute over TLS.
	}

	BootCapabilitiesResponse struct {
		XMLName                            xml.Name `xml:"AMT_BootCapabilities"`
		InstanceID                         string   `xml:"InstanceID,omitempty"`                         // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. In order to ensure uniqueness within the NameSpace, the value of InstanceID SHOULD be constructed using the following 'preferred' algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon ':', and where <OrgID> MUST include a copyrighted, trademarked or otherwise unique name that is owned by the business entity creating/defining the InstanceID, or is a registered ID that is assigned to the business entity by a recognized global authority (This is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness <OrgID> MUST NOT contain a colon (':'). When using this algorithm, the first colon to appear in InstanceID MUST appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and SHOULD not be re-used to identify different underlying (real-world) elements. If the above 'preferred' algorithm is not used, the defining entity MUST assure that the resultant InstanceID is not re-used across any InstanceIDs produced by this or other providers for this instance's NameSpace. For DMTF defined instances, the 'preferred' algorithm MUST be used with the <OrgID> set to 'CIM'.
		ElementName                        string   `xml:"ElementName,omitempty"`                        // The user friendly name for this instance of Capabilities. In addition, the user friendly name can be used as a index property for a search of query. (Note: Name does not have to be unique within a namespace.)
		IDER                               bool     `xml:"IDER,omitempty"`                               // Indicates whether Intel® AMT device supports 'IDE Redirection'
		SOL                                bool     `xml:"SOL,omitempty"`                                // Indicates whether Intel® AMT device supports 'Serial Over Lan'
		BIOSReflash                        bool     `xml:"BIOSReflash,omitempty"`                        // Indicates whether Intel® AMT device supports 'BIOS Reflash'
		BIOSSetup                          bool     `xml:"BIOSSetup,omitempty"`                          // Indicates whether Intel® AMT device supports 'BIOS Setup'
		BIOSPause                          bool     `xml:"BIOSPause,omitempty"`                          // Indicates whether Intel® AMT device supports 'BIOS Pause'
		ForcePXEBoot                       bool     `xml:"ForcePXEBoot,omitempty"`                       // Indicates whether Intel® AMT device supports 'Force PXE Boot'
		ForceHardDriveBoot                 bool     `xml:"ForceHardDriveBoot,omitempty"`                 // Indicates whether Intel® AMT device supports 'Force Hard Drive Boot'
		ForceHardDriveSafeModeBoot         bool     `xml:"ForceHardDriveSafeModeBoot,omitempty"`         // Indicates whether Intel® AMT device supports 'Force Hard Drive Safe Mode Boot'
		ForceDiagnosticBoot                bool     `xml:"ForceDiagnosticBoot,omitempty"`                // Indicates whether Intel® AMT device supports 'Force Diagnostic Boot'
		ForceCDorDVDBoot                   bool     `xml:"ForceCDorDVDBoot,omitempty"`                   // Indicates whether Intel® AMT device supports 'Force CD or DVD Boot'
		VerbosityScreenBlank               bool     `xml:"VerbosityScreenBlank,omitempty"`               // Indicates whether Intel® AMT device supports 'Verbosity Screen Blank'
		PowerButtonLock                    bool     `xml:"PowerButtonLock,omitempty"`                    // Indicates whether Intel® AMT device supports 'Power Button Lock'
		ResetButtonLock                    bool     `xml:"ResetButtonLock,omitempty"`                    // Indicates whether Intel® AMT device supports 'Reset Button Lock'
		KeyboardLock                       bool     `xml:"KeyboardLock,omitempty"`                       // Indicates whether Intel® AMT device supports 'Keyboard Lock'
		SleepButtonLock                    bool     `xml:"SleepButtonLock,omitempty"`                    // Indicates whether Intel® AMT device supports 'Sleep Button Lock'
		UserPasswordBypass                 bool     `xml:"UserPasswordBypass,omitempty"`                 // Indicates whether Intel® AMT device supports 'User Password Bypass'
		ForcedProgressEvents               bool     `xml:"ForcedProgressEvents,omitempty"`               // Indicates whether Intel® AMT device supports 'Forced Progress Events'
		VerbosityVerbose                   bool     `xml:"VerbosityVerbose,omitempty"`                   // Indicates whether Intel® AMT device supports 'Verbosity/Verbose'
		VerbosityQuiet                     bool     `xml:"VerbosityQuiet,omitempty"`                     // Indicates whether Intel® AMT device supports 'Verbosity/Quiet'
		ConfigurationDataReset             bool     `xml:"ConfigurationDataReset,omitempty"`             // Indicates whether Intel® AMT device supports 'Configuration Data Reset'
		BIOSSecureBoot                     bool     `xml:"BIOSSecureBoot,omitempty"`                     // Indicates whether Intel® AMT device supports 'BIOS Secure Boot'
		SecureErase                        bool     `xml:"SecureErase,omitempty"`                        // Indicates whether Intel® AMT device supports 'Secure Erase'
		ForceWinREBoot                     bool     `xml:"ForceWinREBoot,omitempty"`                     // Supports Intel AMT invoking boot to WinRE
		ForceUEFILocalPBABoot              bool     `xml:"ForceUEFILocalPBABoot,omitempty"`              // Supports booting to an ISV’s PBA
		ForceUEFIHTTPSBoot                 bool     `xml:"ForceUEFIHTTPSBoot,omitempty"`                 // Supports Intel AMT invoking HTTPS boot
		AMTSecureBootControl               bool     `xml:"AMTSecureBootControl,omitempty"`               // Determines whether Intel AMT is privileged by BIOS to disable secure boot for an AMT triggered boot option. If true, the BIOS allows Intel AMT to control the secure boot (i.e., to disable secure boot in recovery from HTTPS under certain conditions).
		UEFIWiFiCoExistenceAndProfileShare bool     `xml:"UEFIWiFiCoExistenceAndProfileShare,omitempty"` // Read-only field, determines whether UEFI BIOS and Intel AMT WiFi profile share is supported. The feature is available from Intel® CSME 16.0.
		PlatformErase                      int      `xml:"PlatformErase,omitempty"`                      // Indicates whether the Intel AMT device supports Intel Remote Platform Erase (i.e., whether the OEM's BIOS includes support for the feature), and shows the devices that can be erased. The feature is available from Intel® CSME 16.0.
	}
)

// INPUTS
// Request Types
type BootSettingDataRequest struct {
	XMLName                  xml.Name          `xml:"h:AMT_BootSettingData"`
	H                        string            `xml:"xmlns:h,attr"`
	InstanceID               string            `xml:"h:InstanceID"`               // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>.  <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
	ElementName              string            `xml:"h:ElementName"`              // Required. The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	OwningEntity             string            `xml:"h:OwningEntity"`             // OwningEntity identifies the vendor or organization that defines the contained boot settings.
	UseSOL                   bool              `xml:"h:UseSOL"`                   // Required. When True, Serial over LAN is used on the next boot cycle.
	UseSafeMode              bool              `xml:"h:UseSafeMode"`              // Required. When a Hard-drive boot source is chosen (using CIM_BootConfigSetting) and this property is set to True, the Intel® AMT firmware will boot in safe mode.
	ReflashBIOS              bool              `xml:"h:ReflashBIOS"`              // Required. When True, the Intel® AMT firmware reflashes the BIOS on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
	BIOSSetup                bool              `xml:"h:BIOSSetup"`                // Required. When True, the Intel® AMT firmware enters the CMOS Setup screen on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
	BIOSPause                bool              `xml:"h:BIOSPause"`                // Required. When True, the BIOS pauses for user input on the next boot cycle. This property can be set to true only when a boot source isn't set (using CIM_BootConfigSetting.ChangeBootOrder method).
	LockPowerButton          bool              `xml:"h:LockPowerButton"`          // Required. When True, the Intel® AMT firmware disables the power button operation for the system, normally until the next boot cycle.
	LockResetButton          bool              `xml:"h:LockResetButton"`          // Required. When True, the Intel® AMT firmware disables the reset button operation for the system, normally until the next boot cycle.
	LockKeyboard             bool              `xml:"h:LockKeyboard"`             // Required. When True, the Intel® AMT firmware disallows keyboard activity during its boot process.
	LockSleepButton          bool              `xml:"h:LockSleepButton"`          // Required. When True, the Intel® AMT firmware disables the sleep button operation for the system, normally until the next boot cycle.
	UserPasswordBypass       bool              `xml:"h:UserPasswordBypass"`       // Required. When True, the Intel® AMT firmware boots the system and bypasses any user or boot password that might be set in the system.
	ForcedProgressEvents     bool              `xml:"h:ForcedProgressEvents"`     // Required. When True, the Intel® AMT firmware transmits all progress PET events to the alert-sending device.
	FirmwareVerbosity        FirmwareVerbosity `xml:"h:FirmwareVerbosity"`        // Required. When set to a non-zero value, controls the amount of information the managed system writes to its local display.
	ConfigurationDataReset   bool              `xml:"h:ConfigurationDataReset"`   // Required. When True, the Intel® AMT firmware resets its non-volatile configuration data to the managed system's Setup defaults prior to booting the system.
	IDERBootDevice           IDERBootDevice    `xml:"h:IDERBootDevice"`           // Required. Specifies the device to use when UseIder is set. 0 - Floppy Boot, 1- CD Boot.
	UseIDER                  bool              `xml:"h:UseIDER"`                  // Required. When True, IDER is used on the next boot cycle.
	EnforceSecureBoot        bool              `xml:"h:EnforceSecureBoot"`        // When True, Secure boot over IDER is enforced on the next boot cycle, if IDER boot is used. This field is also used in One-Click Recovery.
	BootMediaIndex           int               `xml:"h:BootMediaIndex"`           // Required. This property identifies the boot-media index for the managed client (when a boot source is set using the CIM_BootConfigSetting.ChangeBootOrder method). For Hard-Drive or CD/DVD boot - when the parameter value is 0, the default boot-media is booted. When the parameter value is 1, the primary boot-media is booted; when the value is 2, the secondary boot-media is booted; and so on. For PXE or diagnostics boot this property must be 0.
	SecureErase              bool              `xml:"h:SecureErase"`              // Required. When True, the BIOS performs secure erase operation. Note: Customers are recommended to use Secure Remote Platform Erase which is newer and more advanced than this function.
	RSEPassword              string            `xml:"h:RSEPassword"`              // SSD password for Remote Secure Erase operation. This is a write-only field, an empty string is returned when instance is read. When writing, an empty string or lack of field will be ignored. The password length is limited to 32 ASCII characters. Note: Customers are recommended to use Secure Remote Platform Erase which is newer and more advanced than Remote Secure Erase.
	OptionsCleared           bool              `xml:"h:OptionsCleared"`           // Indicates whether the boot options have been cleared by BIOS or not. This property is read only.
	WinREBootEnabled         bool              `xml:"h:WinREBootEnabled"`         // Indicates whether ForceWinREBoot is enabled in BIOS. This property is read only.
	UEFILocalPBABootEnabled  bool              `xml:"h:UEFILocalPBABootEnabled"`  // Indicates whether ForceUEFILocalPBABoot is enabled in BIOS. This property is read only.
	UEFIHTTPSBootEnabled     bool              `xml:"h:UEFIHTTPSBootEnabled"`     // Indicates whether ForceUEFIHTTPSBoot is enabled in BIOS. This property is read only.
	SecureBootControlEnabled bool              `xml:"h:SecureBootControlEnabled"` // Determines whether Intel AMT is privileged by BIOS to disable secure boot for an AMT triggered boot option. If not, BIOSSecureBoot must be set to TRUE. This property is read only.
	BootguardStatus          bool              `xml:"h:BootguardStatus"`          // Enables the console to discover the security level of the BIOS boot flow. This property is read only.
	BIOSLastStatus           []int             `xml:"h:BIOSLastStatus"`           // Last boot status reported by BIOS. The first 16-bit word contains the general BIOS status (0 - Success, 1 - In Progress, 2 - Not Updated, 0xFFFF - Failed). The second word contains the detailed error status (0 - Success/In Progress, 1 - General Drive Failure, 2 - Drive Password/Authentication Failure, 3 - Feature is not supported). This property is read-only.
	UEFIBootParametersArray  []int             `xml:"h:UEFIBootParametersArray"`  // TLV parameters array encoded with base64 for configuring boot parameters for One-Click Recovery and Secure Remote Platform Erase.
	UEFIBootNumberOfParams   []int             `xml:"h:UEFIBootNumberOfParams"`   // Number of parameters in UefiBootParametersArray
	RPEEnabled               bool              `xml:"h:RPEEnabled"`               // Indicates whether Secure Remote Platform Erase is enabled by the BIOS. Note: This command needs to execute over TLS.
	PlatformErase            bool              `xml:"h:PlatformErase"`            // When set to True, sets the boot option to trigger Secure Remote Platform Erase in the next boot.  Note: This command needs to execute over TLS.
}

// When set to a non-zero value, controls the amount of information the managed system writes to its local display.
//
// ValueMap={0, 1, 2, 3}
//
// Values={System default, Quiet - minimal screen activity, Verbose - all messages appear on the screen, Screen blank - no messages appear on the screen}
type FirmwareVerbosity int

// Specifies the device to use when UseIder is set. 0 - Floppy Boot, 1- CD Boot.
//
// ValueMap={0, 1}
//
// Values={Floppy Boot, CD Boot}
type IDERBootDevice int
