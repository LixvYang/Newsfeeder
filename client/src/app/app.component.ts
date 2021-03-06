import { Component, OnInit } from '@angular/core';
import { HttpClient} from '@angular/common/http'
interface INewsfeedItem {
  title: string
  post: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = ''
  public post = ''
  public newsfeedItems:INewsfeedItem[] = []

  constructor(
    private httpClient:HttpClient 
  ){}

  async ngOnInit(){
    await this.loadNewsItems()

    setInterval(() => this.loadNewsItems(),2000)
  }



  async loadNewsItems(){
    this.newsfeedItems = await this.httpClient
    .get<INewsfeedItem[]>('/api/newsfeed')
    .toPromise()
  }

  async addPost(){
    await this.httpClient.post('./api/newsfeed',{
      title:this.title,
      post:this.post
    }).toPromise()

    await this.loadNewsItems()

    this.title = ''
    this.post = ''
  }
}
