import { Component } from '@angular/core';
import { first } from 'rxjs/operators';
import { User } from '../_models';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { NewPostComponent } from '../tools'
import { UserService, AuthenticationService, RestService } from '../_services';



@Component({ templateUrl: 'home.component.html' })
export class HomeComponent {
  user: User
  posts: any[];

  constructor(
    private userService: UserService,
    private restService: RestService,
    private authService: AuthenticationService,
    private modalService: NgbModal
  ) { }

  ngOnInit() {
    this.user = this.authService.currentUserValue;
    this.restService.timeline().subscribe(response => {
      this.posts = response.posts;
    })

    // this.userService.getAll().pipe(first()).subscribe(users => {
    //     this.users = users;
    // });
  }

  open() {
    const modalRef = this.modalService.open(NewPostComponent);
  }
}
