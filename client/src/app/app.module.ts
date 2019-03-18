import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';


import { JwtInterceptor, ErrorInterceptor } from './_helpers';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { HomeComponent } from './home';
import { LoginComponent } from './login';
import { NewPostComponent } from './tools'

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    NewPostComponent
  ],
  entryComponents: [NewPostComponent],
  imports: [
    BrowserModule,
    ReactiveFormsModule,
    HttpClientModule,
    NgbModule,
    AppRoutingModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS,
      useClass: JwtInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS,
      useClass: ErrorInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
