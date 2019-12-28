import React, { Component } from "react";
import SelectedTags from "./SelectedTags";
import TagSearch from "./TagSearch";
import PictureResult from "./PictureResult"
import Client from "./Client";
import { Button } from 'semantic-ui-react';


class App extends Component {
  state = {
    selectedTags: [],
    selectedPictures: []
  };

  removeTagItem = itemIndex => {
    const filteredTags = this.state.selectedTags.filter(
      (item, idx) => itemIndex !== idx
    );
    this.setState({ selectedTags: filteredTags });
  };

  addTag = tag => {
    var bool = 0
    for(var i=0; i<this.state.selectedTags.length; i++){
      if (this.state.selectedTags[i].id === tag.id) {
        bool = 1
      }
    }
    if (bool === 0)
    {
      const newTags = this.state.selectedTags.concat(tag);
      this.setState({ selectedTags: newTags });
    }
  };

  showPictures =  async picture => {
    let all_pictures = []
    for(var i=0; i<this.state.selectedTags.length; i++){
      console.log(this.state.selectedTags[i].id.substring(3))
      all_pictures.push(await Client.search(this.state.selectedTags[i].id.substring(3)));
    }
    console.log(all_pictures)
    this.setState({
      selectedPictures: all_pictures[0].Picture
    });
  };

  render() { 
    const { selectedTags, selectedPictures } = this.state;

    return (
      <div className="App">
        <div className="ui text container">
          <SelectedTags
            tags={selectedTags}
            onTagClick={this.removeTagItem}
          />
          <TagSearch onTagClick={this.addTag} />
          <Button onClick={this.showPictures}>
              Result
          </Button>
        </div>
          <PictureResult
            pictures={selectedPictures}
          />
      </div>
    );
  }
}

export default App;
