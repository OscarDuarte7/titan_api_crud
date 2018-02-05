package routers

import (
	"github.com/astaxie/beego"
)

func init() {
/*
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method:           "GetConceptosSs",
			Router:           `get_conceptos_ss/:persona`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiariosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:BeneficiariosController"],
		beego.ControllerComments{
			Method: "BeneficiarioDatos",
			Router: `/beneficiarioDatos`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})
*/

beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
	beego.ControllerComments{
		Method:           "ConsultarAsignacionBasica",
		Router:           `/`,
		AllowHTTPMethods: []string{"post"},
		Params:           nil})

beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:DocenteCargoController"],
	beego.ControllerComments{
		Method:           "ConsultarCedulaDocente",
		Router:           `/consultarCedulaDocente`,
		AllowHTTPMethods: []string{"post"},
		Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioProveedorController"],
		beego.ControllerComments{
			Method:           "ConsultarIDProveedor",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

		beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioPritecController"],
					beego.ControllerComments{
					Method:           "ConsultarPrimaTecnica",
					Router:           `/`,
					AllowHTTPMethods: []string{"post"},
					Params:           nil})

		beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:FuncionarioCargoController"],
					beego.ControllerComments{
					Method:           "ConsultarAsignacionBasica",
					Router:           `/`,
					AllowHTTPMethods: []string{"post"},
					Params:           nil})


	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:PreliquidacionController"],
		beego.ControllerComments{
			Method:           "Resumen",
			Router:           `/resumen`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})


		/*
	beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LiquidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:LiquidacionController"],
		beego.ControllerComments{
			Method:           "Resumen",
			Router:           `/resumen`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})
		*/


/*
				beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:InformacionPensionadoController"],
					beego.ControllerComments{
						Method: "Pensionado_datos",
						Router: `/`,
						AllowHTTPMethods: []string{"post"},
						Params: nil})

					beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
						beego.ControllerComments{
							Method: "Sustituto_datos",
							Router: `/sustitutoDatos`,
							AllowHTTPMethods: []string{"post"},
							Params: nil})

						beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/titan_api_crud/controllers:SustitutoController"],
							beego.ControllerComments{
								Method: "Tutor_datos",
								Router: `/tutorDatos`,
								AllowHTTPMethods: []string{"post"},
								Params: nil})
								*/
}
