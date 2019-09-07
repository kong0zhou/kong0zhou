import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { ShowComponent } from './components/show/show.component';
import { HttpClientModule } from '@angular/common/http';
import {
  MatTableModule, MatButtonModule, MatPaginatorModule,
  MatCardModule, MatDialogModule, MatSnackBarModule,
  MatProgressSpinnerModule, MatSelectModule, MatTabsModule,
  MatExpansionModule, MatFormFieldModule, MatInputModule,
  MatIconModule,MatRippleModule,MatDatepickerModule,MatNativeDateModule,MatSidenavModule
} from '@angular/material';

@NgModule({
  declarations: [
    AppComponent,
    ShowComponent
  ],
  imports: [
    MatButtonModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSidenavModule,
    HttpClientModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
