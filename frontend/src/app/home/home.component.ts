import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  name: string | null = '';
  userId: string | null = '';
  favColor: string | null = '#FFFFFF';

  constructor(
    private router: Router
  ) { 
    // navigation is an object that has info about the current navigation action, ie url, params, states
    // in register component, i pass the user's name and user id so it can be dispalyed
    const navigation = this.router.getCurrentNavigation();
    // extract the state from the current navigation
    const state = navigation?.extras.state as { name: string, userId: string, favColor: string };
    this.favColor = state.favColor;

    if (state) {
      this.name = state.name;
      this.userId = state.userId;
    }
    else {
      this.name = "guest";
      this.userId = "-1";
    }
  }

  ngOnInit(): void {
  }
}
