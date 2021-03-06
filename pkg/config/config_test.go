package config

import (
	"reflect"
	"testing"
)

const (
	configFilePath                   = "../../configs/testConfigs/test-config.yaml"
	correctTestConfigName            = "UptimeRobot"
	correctTestAPIKey                = "657a68d9ashdyasjdklkskuasd"
	correctTestAPIURL                = "https://api.uptimerobot.com/v2/"
	correctTestAlertContacts         = "0544483_0_0-2628365_0_0-2633263_0_0"
	correctTestEnableMonitorDeletion = true

	configFilePathPingdom          = "../../configs/testConfigs/test-config-pingdom.yaml"
	correctTestPingdomConfigMulti  = "PingdomMulti"
	correctTestPingdomUsername     = "user@test.com"
	correctTestPingdomAPIURL       = "https://api.pingdom.com/v2/"
	correctTestPingdomPassword     = "SuperSecret"
	correctTestPingdomAccountEmail = "multi@test.com"
)

func TestConfigWithCorrectValues(t *testing.T) {
	correctConfig := Config{Providers: []Provider{Provider{Name: correctTestConfigName, ApiKey: correctTestAPIKey, ApiURL: correctTestAPIURL, AlertContacts: correctTestAlertContacts}}, EnableMonitorDeletion: correctTestEnableMonitorDeletion}
	config := ReadConfig(configFilePath)

	if !reflect.DeepEqual(config, correctConfig) {
		t.Error("Marshalled config and correct config do not match")
	}
}

func TestConfigWithIncorrectProviderValues(t *testing.T) {
	incorrectConifg := Config{Providers: []Provider{Provider{Name: "UptimeRobot2", ApiKey: "abc", ApiURL: correctTestAPIURL, AlertContacts: correctTestAlertContacts}}, EnableMonitorDeletion: correctTestEnableMonitorDeletion}
	config := ReadConfig(configFilePath)

	if reflect.DeepEqual(config, incorrectConifg) {
		t.Error("Marshalled config and incorrect config match, should not match")
	}
}

func TestConfigWithIncorrectEnableFlag(t *testing.T) {
	incorrectConifg := Config{Providers: []Provider{Provider{Name: correctTestConfigName, ApiKey: correctTestAPIKey, ApiURL: correctTestAPIURL, AlertContacts: correctTestAlertContacts}}, EnableMonitorDeletion: false}
	config := ReadConfig(configFilePath)

	if reflect.DeepEqual(config, incorrectConifg) {
		t.Error("Marshalled config and incorrect config match, should not match")
	}
}

func TestConfigWithoutProvider(t *testing.T) {
	incorrectConifg := Config{Providers: []Provider{}, EnableMonitorDeletion: false}
	config := ReadConfig(configFilePath)

	if reflect.DeepEqual(config, incorrectConifg) {
		t.Error("Marshalled config and incorrect config match, should not match")
	}
}

func TestConfigWithoutEnabledFlag(t *testing.T) {
	incorrectConifg := Config{Providers: []Provider{Provider{Name: correctTestConfigName, ApiKey: correctTestAPIKey, ApiURL: correctTestAPIURL, AlertContacts: correctTestAlertContacts}}}
	config := ReadConfig(configFilePath)

	if reflect.DeepEqual(config, incorrectConifg) {
		t.Error("Marshalled config and incorrect config match, should not match")
	}
}

func TestConfigWithPingdomMultiAuthEnabledFlag(t *testing.T) {
	correctConfig := Config{Providers: []Provider{Provider{Name: correctTestPingdomConfigMulti, ApiKey: correctTestAPIKey, ApiURL: correctTestPingdomAPIURL, AlertContacts: correctTestAlertContacts, Username: correctTestPingdomUsername, Password: correctTestPingdomPassword, AccountEmail: correctTestPingdomAccountEmail}}, EnableMonitorDeletion: correctTestEnableMonitorDeletion}
	config := ReadConfig(configFilePathPingdom)
	if !reflect.DeepEqual(config, correctConfig) {
		t.Error("Marshalled config and correct config do not match")
	}
}

func TestConfigWithEmptyConfig(t *testing.T) {
	incorrectConifg := Config{}
	config := ReadConfig(configFilePath)

	if reflect.DeepEqual(config, incorrectConifg) {
		t.Error("Marshalled config and incorrect config match, should not match")
	}
}
