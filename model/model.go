package model

type OneField struct {
	Field1 string `json:"field1"`
}

type TwoFields struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

type Indicator struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Label string `json:"label"`
}

type AllFields struct {
	D_DPTO            string `json:"D_DPTO"`
	D_PROV            string `json:"D_PROV"`
	D_DIST            string `json:"D_DIST"`
	D_DREUGEL         string `json:"D_DREUGEL"`
	RURAL_PMM         string `json:"RURAL_PMM"`
	RURAL_PMM_MUJE1   string `json:"RURAL_PMM_MUJE1"`
	RURAL_PMM_MUJE2   string `json:"RURAL_PMM_MUJE2"`
	RURAL_PMM_HOME1   string `json:"RURAL_PMM_HOME1"`
	RURAL_PMM_HOME2   string `json:"RURAL_PMM_HOME2"`
	RURAL_PMMA_MUJE1  string `json:"RURAL_PMMA_MUJE1"`
	RURAL_PMMA_MUJE2  string `json:"RURAL_PMMA_MUJE2"`
	RURAL_PMMA_HOME1  string `json:"RURAL_PMMA_HOME1"`
	RURAL_PMMA_HOME2  string `json:"RURAL_PMMA_HOME2"`
	RURAL_PMM_MUJDOC  string `json:"RURAL_PMM_MUJDOC"`
	RURAL_PMM_HOMDOC  string `json:"RURAL_PMM_HOMDOC"`
	RURAL_PMMA_MUJDOC string `json:"RURAL_PMMA_MUJDOC"`
	RURAL_PMMA_HOMDOC string `json:"RURAL_PMMA_HOMDOC"`
	RURAL_CRFA        string `json:"RURAL_CRFA"`
	RURAL_SRE         string `json:"RURAL_SRE"`
	RURAL_ST          string `json:"RURAL_ST"`
	RURAL_CRFA_MUJE1  string `json:"RURAL_CRFA_MUJE1"`
	RURAL_CRFA_MUJE2  string `json:"RURAL_CRFA_MUJE2"`
	RURAL_CRFA_HOME1  string `json:"RURAL_CRFA_HOME1"`
	RURAL_CRFA_HOME2  string `json:"RURAL_CRFA_HOME2"`
	RURAL_SRE_MUJE1   string `json:"RURAL_SRE_MUJE1"`
	RURAL_SRE_MUJE2   string `json:"RURAL_SRE_MUJE2"`
	RURAL_SRE_HOME1   string `json:"RURAL_SRE_HOME1"`
	RURAL_SRE_HOME2   string `json:"RURAL_SRE_HOME2"`
	RURAL_MSE_MUJDOC  string `json:"RURAL_MSE_MUJDOC"`
	RURAL_MSE_HOMDOC  string `json:"RURAL_MSE_HOMDOC"`
}

type Tuple [2]int

type Group struct {
	NAME string  `json:"name"`
	DATA [][]int `json:"data"`
}
