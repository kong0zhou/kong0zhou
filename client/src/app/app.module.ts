import { BrowserModule } from '@angular/platform-browser';
import { NgModule,Optional, SkipSelf } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { ShowComponent } from './components/show/show.component';

// >>>>>>>>> 加载icon库 >>>>>>>>>>
import { loadSvgsources } from 'src/app/icon.svg';
import { MatIconRegistry } from '@angular/material';
import { DomSanitizer } from '@angular/platform-browser';
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

import { HttpClientModule } from '@angular/common/http';
import {
  MatTableModule, MatButtonModule, MatPaginatorModule,
  MatCardModule, MatDialogModule, MatSnackBarModule,
  MatProgressSpinnerModule, MatSelectModule, MatTabsModule,
  MatExpansionModule, MatFormFieldModule, MatInputModule,
  MatIconModule,MatRippleModule,MatDatepickerModule,MatNativeDateModule,
  MatSidenavModule,MatTreeModule,MatButtonToggleModule,
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
    MatTreeModule,
    MatIconModule,
    // MatButtonToggleModule,
    // MatInputModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { 
  constructor(
    @Optional() @SkipSelf() parent: AppModule,
    ir: MatIconRegistry,
    ds: DomSanitizer
  ) {
    if (parent) {
      throw new Error('模块已经存在，不能再次加载');
    }
    loadSvgsources(ir, ds);
  }
}
