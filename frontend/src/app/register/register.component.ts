import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  constructor(
    private userService: UserService,
    private router: Router,
  ) { }

  ngOnInit(): void { }

  isValidHexColor(color: string): boolean {
    const hexColorRegex = /^#(?:[0-9a-fA-F]{3}){1,2}$/;
    return hexColorRegex.test(color);
  }

  isValidEmail(email: string): boolean {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
  }

  isValidName(name: string): boolean {
    return name.length > 0;
  }

  onFormSubmit(form: NgForm) {
    if (this.isValidName(form.value.name) && this.isValidEmail(form.value.email)) {
      const name = form.value.name;
      const email = form.value.email;
      const favColor = this.isValidHexColor(form.value.favColor) ? form.value.favColor : '#FFFFFF';

      this.userService.postRegister(name, email, favColor).subscribe((response) => {
        // Once we've received a response, take the user to the home page
        // store the user's name and user id for the home page
        const navigationExtras = {
          state: {
            name: name,
            userId: response.user_id,
            favColor: favColor
          }
        };
        this.router.navigateByUrl('/home', navigationExtras);
      })
    }
    else {
      // handle invalid input
      alert("Invalid input. Please enter a valid name and email address");
    }
  }
}
