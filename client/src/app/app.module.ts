import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';


import { JwtInterceptor, ErrorInterceptor } from './_helpers';
import { HomeComponent } from './home';
import { LoginComponent } from './login';

@NgModule({
    declarations: [
        AppComponent,
        HomeComponent,
        LoginComponent
    ],
    imports: [
        BrowserModule,
        ReactiveFormsModule,
        HttpClientModule,
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
export class AppModule { }
