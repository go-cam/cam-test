import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { FileRoutingModule } from './file-routing.module';
import { FileComponent } from './file.component';
import { FileUploader, FileUploadModule } from 'ng2-file-upload';


@NgModule({
  declarations: [FileComponent],
  imports: [
    CommonModule,
    FileRoutingModule,
    FileUploadModule
  ]
})
export class FileModule {

}
