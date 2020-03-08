import { Component, OnInit } from '@angular/core';
import { FileUploader } from 'ng2-file-upload';

@Component({
  selector: 'app-file',
  templateUrl: './file.component.html',
  styleUrls: ['./file.component.scss']
})
export class FileComponent implements OnInit {
  public uploader: FileUploader;

  constructor() {
    this.uploader = new FileUploader({url: "/file/upload"})
  }

  ngOnInit(): void {
  }

}
