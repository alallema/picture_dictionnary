import React, { Component } from "react";
import SelectedTags from "./SelectedTags";
import TagSearch from "./TagSearch";
import PictureResult from "./PictureResult"
import Client from "./Client";
import { Container, Button, Divider } from 'semantic-ui-react';

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
    this.setState({
      selectedPictures: []
    });
    var result = Array.from(this.state.selectedTags, x => x.id.replace(/\//g, ''))
    let all_pictures = []
    if (result.length !== 0)
      all_pictures = await Client.search(result.join(','))
    this.setState({
      selectedPictures: all_pictures.result
    });
  };

  render() { 
    const { selectedTags, selectedPictures } = this.state;

    return (
      <Container style={{ marginTop: '3em' }}>
        <div className="App">
          <SelectedTags
            tags={selectedTags}
            onTagClick={this.removeTagItem}
          />
          <Divider />
          <TagSearch onTagClick={this.addTag} />
          <Divider />
          <Button onClick={this.showPictures}>
              Result
          </Button>
          <Divider />
          <PictureResult
            pictures={selectedPictures}
          />
        </div>
      </Container>
    );
  }
}

export default App;
