import React from "react";
import { Icon, Label } from 'semantic-ui-react';

// const colors = [
//   'red',
//   'orange',
//   'yellow',
//   'olive',
//   'green',
//   'teal',
//   'blue',
//   'violet',
//   'purple',
//   'pink',
//   'brown',
//   'grey',
//   'black',
// ]

export default function SelectedTags(props) {
  const { tags } = props;

  const tagRows = tags.map((tag, idx) => (
      <Label key={idx} style={{ marginBottom: '1em' }} onClick={() => props.onTagClick(idx)}>
        {tag.title}
        <Icon name='close' />
      </Label>
  ));

  return (
      <div>
        {tagRows}
      </div>
  );
}