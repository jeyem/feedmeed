import { Component } from '@angular/core';
import { first } from 'rxjs/operators';

import { User } from '../_models';
import { UserService, AuthenticationService } from '../_services';

@Component({ templateUrl: 'home.component.html' })
export class HomeComponent {
  user: User

  constructor(
    private userService: UserService,
    private authService: AuthenticationService
  ) { }

  ngOnInit() {
    this.user = this.authService.currentUserValue;
    // this.userService.getAll().pipe(first()).subscribe(users => {
    //     this.users = users;
    // });
  }
}
