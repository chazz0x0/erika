package erika

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type marshalable interface{}

func getGenericAPICall(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getGenericUnmarshalledAPICall(url string, m marshalable) error {
	body, err := getGenericAPICall(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}

	return nil
}

// Secret Keys
func GetSecretKey(version string) (string, error) {
	// TODO: This needs improved. Severely.
	resp, err := getGenericAPICall("https://api.wordpress.org/secret-key/" + version + "/")
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// Stats calls
func getGenericVersionStats(url string) (map[string]float64, error) {
	results := make(map[string]float64)
	err := getGenericUnmarshalledAPICall(url, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetWordPressVersionStats returns a map of WordPress version numbers to their
// usage ratio.
func GetWordPressVersionStats() (map[string]float64, error) {
	return getGenericVersionStats("https://api.wordpress.org/stats/wordpress/1.0/")
}

// GetWordPressVersionStats returns a map of PHP version numbers to their
// usage ratio.
func GetPHPVersionStats() (map[string]float64, error) {
	return getGenericVersionStats("https://api.wordpress.org/stats/php/1.0/")
}

// GetWordPressVersionStats returns a map of MySQL/MariaDB version numbers to
// their usage ratio.
func GetMySQLVersionStats() (map[string]float64, error) {
	return getGenericVersionStats("https://api.wordpress.org/stats/mysql/1.0/")
}

func GetLocaleVersionStats() (map[string]float64, error) {
	return getGenericVersionStats("https://api.wordpress.org/stats/locale/1.0/")
}

func GetPluginVersionStats(slug string) (map[string]float64, error) {
	return getGenericVersionStats("https://api.wordpress.org/stats/plugin//1.0/" + slug)
}

func GetPluginDownloadStats(slug string, limit int) (map[string]float64, error) {
	// TODO: Figure out what this one does.
	// https://api.wordpress.org/stats/plugin/1.0/downloads.php?slug={slug}&limit={days}&callback={jsFunction}
	return nil, nil
}

// Version Check

// TODO: Is this even worth implementing?
func VersionCheck() {
	//return
}

// Version Stability
func GetVersionStabilityList() (map[string]string, error) {
	m := make(map[string]string)
	err := getGenericUnmarshalledAPICall("https://api.wordpress.org/core/stable-check/1.0/", &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Credits
func GetCredits(version string, locale string) (*Credits, error) {
	var credits Credits
	err := getGenericUnmarshalledAPICall(
		fmt.Sprintf("https://api.wordpress.org/core/credits/1.1/?version=%s&locale=%s", version, locale),
		&credits,
	)
	if err != nil {
		return nil, err
	}

	return &credits, nil
}

// Plugins

// TODO: This needs arguments added, but works for my usecase for now.
func SearchPlugins() (*PluginSearch, error) {
	var res PluginSearch
	err := getGenericUnmarshalledAPICall(
		"https://api.wordpress.org/plugins/info/1.2/?action=query_plugins&browse=popular",
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
