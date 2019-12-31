import React, { Component } from 'react'
import Client from "./Client";
import _ from 'lodash';
import { Grid, Search } from 'semantic-ui-react';

const initialState = { isLoading: false, results: [], value: '' }

export default class TagSearch extends Component {
  state = initialState

  handleSearchChange = (e, { value }) => {
    this.setState({
        isLoading: true, value
    })

    setTimeout(async () => {
    if (this.state.value.length < 1) return this.setState(initialState)

    const re = new RegExp(_.escapeRegExp(this.state.value), 'i')
    const isMatch = (result) => re.test(result.title)

    var results = await Client.translate(value)
    this.setState({
        isLoading: false,
        results: _.filter(results.result, isMatch),
      })
    }, 30)
  }

  render() {
    const { isLoading, value, results } = this.state

    return (
      <Grid>
        <Grid.Column width={6}>
          <Search
            loading={isLoading}
            onResultSelect={(e, {result}) => {
              this.props.onTagClick(result)
            }}
            onSearchChange={_.debounce(this.handleSearchChange, 300, {
              leading: true,
            })}
            results={results}
            value={value}
          />
        </Grid.Column>
        {/* <Grid.Column width={10}>
          <Segment>
            <Header>State</Header>
            <pre style={{ overflowX: 'auto' }}>
              {JSON.stringify(this.state, null, 2)}
            </pre>
            <Header>Options</Header>
            <pre style={{ overflowX: 'auto' }}>
              {JSON.stringify(this.results, null, 2)}
            </pre>
          </Segment>
        </Grid.Column> */}
      </Grid>
    )
  }
}
