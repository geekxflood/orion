// internal/modules/glpi/glpi.go

package module_glpi

func GetGLPI() error {

	// Full doc: https://github.com/ramylson/glpi/blob/master/glpi/apirest.md

	// Init session for getting session token
	/*

		URL: apirest.php/initSession/

		Description: Request a session token to uses other api endpoints.

		Method: GET

		Parameters: (Headers)

		App-Token: authorization string provided by the GLPI api configuration. Optional.

		a couple login & password: 2 parameters to login with user authentication. You should pass this 2 parameters in http basic auth. It consists in a Base64 string with login and password separated by ":" A valid Authorization header is: * "Authorization: Basic base64({login}:{password})"

		OR

		an user_token defined in User Preference (See 'Remote access key') You should pass this parameter in 'Authorization' HTTP header. A valid Authorization header is: * "Authorization: user_token q56hqkniwot8wntb3z1qarka5atf365taaa2uyjrn"

		Returns:

		200 (OK) with the session_token string.
		400 (Bad Request) with a message indicating an error in input parameter.
		401 (UNAUTHORIZED)
	*/

	return nil
}
