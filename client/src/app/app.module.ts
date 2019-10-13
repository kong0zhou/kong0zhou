import { BrowserModule } from '@angular/platform-browser';
import { NgModule,Optional, SkipSelf } from '@angular/core';

import { AppComponent } from './app.component';
import { ShowComponent } from './components/show/show.component';
import { LoginComponent } from './components/login/login.component';

// >>>>>>>>> 加载icon库 >>>>>>>>>>
import { loadSvgsources } from 'src/app/icon.svg';
import { MatIconRegistry } from '@angular/material';
import { DomSanitizer } from '@angular/platform-browser';
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// >>>>>>>>>>>module>>>>>>>>>>>>
import { HttpClientModule } from '@angular/common/http';
import { FormsModule,ReactiveFormsModule  } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { ToastrModule } from 'ngx-toastr';
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// >>>>>>>>> material >>>>>>>>>>
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {
  MatTableModule, MatButtonModule, MatPaginatorModule,
  MatCardModule, MatDialogModule, MatSnackBarModule,
  MatProgressSpinnerModule, MatSelectModule, MatTabsModule,
  MatExpansionModule, MatFormFieldModule, MatInputModule,
  MatIconModule,MatRippleModule,MatDatepickerModule,MatNativeDateModule,
  MatSidenavModule,MatTreeModule,MatButtonToggleModule,
} from '@angular/material';
import {ErrorStateMatcher,ShowOnDirtyErrorStateMatcher} from '@angular/material/core';
//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<,,

import { LogHighLightPipe } from './pipe/log-high-light.pipe';

@NgModule({
  declarations: [
    AppComponent,
    ShowComponent,
    LogHighLightPipe,
    LoginComponent
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
    MatFormFieldModule,
    FormsModule,
    MatInputModule,
    ReactiveFormsModule ,
    ToastrModule.forRoot({
      timeOut: 20000,
      positionClass: 'toast-top-center'
    }),
    // MatTableModule,
  ],
  providers: [
    {provide: ErrorStateMatcher, useClass: ShowOnDirtyErrorStateMatcher},
  ],
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
