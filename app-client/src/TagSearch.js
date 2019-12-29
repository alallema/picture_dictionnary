import React from "react";
import Client from "./Client";

const MATCHING_ITEM_LIMIT = 25;

class TagSearch extends React.Component {
  state = {
    tags: [],
    showRemoveIcon: false,
    searchValue: ""
  };

  handleSearchChange = e => {
    const value = e.target.value;

    this.setState({
      searchValue: value
    });

    if (value === "") {
      this.setState({
        tags: [],
        showRemoveIcon: false
      });
    } else {
      this.setState({
        showRemoveIcon: true
      });

      Client.translate(value, tags => {
        this.setState({
          tags: tags.slice(0, MATCHING_ITEM_LIMIT)
        });
      });
    }
  };

  handleSearchCancel = () => {
    this.setState({
      tags: [],
      showRemoveIcon: false,
      searchValue: ""
    });
  };

  render() {
    const { showRemoveIcon, tags } = this.state;
    const removeIconStyle = showRemoveIcon ? {} : { visibility: "hidden" };

    const tagRows = tags.map((tag, idx) => (
      <tr key={idx} onClick={() => this.props.onTagClick(tag)}>
        <td>{tag.title}</td>
      </tr>
    ));

    return (
      <div id="tag-search">
        <table className="ui selectable structured large table">
          <thead>
            <tr>
              <th colSpan="5">
                <div className="ui fluid search">
                  <div className="ui icon input">
                    <input
                      className="prompt"
                      type="text"
                      placeholder="Search tags..."
                      value={this.state.searchValue}
                      onChange={this.handleSearchChange}
                    />
                    <i className="search icon" />
                  </div>
                  <i
                    className="remove icon"
                    onClick={this.handleSearchCancel}
                    style={removeIconStyle}
                  />
                </div>
              </th>
            </tr>
          </thead>
          <tbody>
            {tagRows}
          </tbody>
        </table>
      </div>
    );
  }
}

export default TagSearch;
