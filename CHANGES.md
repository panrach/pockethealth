## Change 2
ctrl-shift-f 8080 to find all old instances. replaced them all with 80.

modified:
- `LISTENING_PORT` in `main.go`
- `environment.apiUrl` in `environment.prod.ts`
- `environment.apiUrl` in `environment.ts`

## Change 3
I ctrl-shift-f `PostRegister` to find the function. Saw that it was defined in 
`service.go`. I changed the return to userId.

## Change 4
I googled how to make a new page in angular and used the angular CLI to create a new home component. I copied the register page as a template. In `app-routing.module.ts`, I added the HomeComponent. I then googled how to pass data between different components. Here are the changes I made:
- `register.component.ts`: I added a state, which stores user and userId to navigationExtras. I pass this to navigateByUrl when redirecting to the homepage.
- In `home.component.ts`, I extract this data from the navigation object
- In `home.component.html`, I display the extracted data

## Change 5
In the frontend:
- Firstly, I modified `register.component.html` to take favorite color as an input to the form.
- I then went to `register.component.ts` to look for all instances of name and email to add favorite color as form input. I also added favorite color to the state object.
- In `home.component.ts`, I extract the favorite color from the state.

In the backend:
- I modified `users.go` to store the favorite color
- In `controller.go`, I looked for all instances of name and email to store favorite color as well. This includes modifying `PostRegisterBody`, and adding favorite color as a parameter to `PostRegister`.
- I add favorite color to the user details and body in `user.service.ts`
- Now `PostRegister` has an error because there is an extra parameter. I went to `PostRegister` in `service.go` to add favorite color as a parameter to the datastore. 
- Now I can finally display the favorite color. In `home.component.html`, I add favColor as the background.

## Input Sanitization
In the frontend, I do my sanitization in `register.component.ts`
- I created 3 functions to validate the form data. These are `isValidHexColor`, `isValidEmail`, `isValidName`.
- In `onFormSubmit`, I check that these fields are properly filled. If not, I display an error popup.

In the backend, I do my sanitization in `controller.go`
- I created 3 functions to validate name, email, and hex code. These are `ValidateName`, `ValidateEmail`, and `ValidateHexColor`.
- When I receive a request, I validate the data. If any validations fail, I return StatusBadRequest.

## Sources
- https://stackoverflow.com/questions/68191099/how-to-pass-data-between-routed-components-in-angular (for change 4)
