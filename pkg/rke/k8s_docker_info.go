package rke

func loadK8sVersionDockerInfo() map[string][]string {
	return map[string][]string{
		"1.8":  {"1.11.x", "1.12.x", "1.13.x", "17.03.x"},
		"1.9":  {"1.11.x", "1.12.x", "1.13.x", "17.03.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.10": {"1.11.x", "1.12.x", "1.13.x", "17.03.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.11": {"1.11.x", "1.12.x", "1.13.x", "17.03.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.12": {"1.11.x", "1.12.x", "1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.13": {"1.11.x", "1.12.x", "1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.14": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.15": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.16": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.17": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.18": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.19": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.20": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.21": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.22": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x"},
		"1.23": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x"},
		"1.24": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x"},
		"1.25": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x"},
		"1.26": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x"},
		"1.27": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x", "26.0.x", "26.1.x"},
		"1.28": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x", "26.0.x", "26.1.x"},
		"1.29": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x", "26.0.x", "26.1.x"},
		"1.30": {"1.13.x", "17.03.x", "17.06.x", "17.09.x", "18.06.x", "18.09.x", "19.03.x", "20.10.x", "23.0.x", "24.0.x", "25.0.x", "26.0.x", "26.1.x"},
	}
}
