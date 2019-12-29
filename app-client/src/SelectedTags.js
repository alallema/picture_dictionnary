import React from "react";

export default function SelectedTags(props) {
  const { tags } = props;

  const tagRows = tags.map((tag, idx) => (
    <tr key={idx} onClick={() => props.onTagClick(idx)}>
      <td>{tag.description}</td>
      <td className="right aligned">{tag.title}</td>
    </tr>
  ));

  return (
    <table className="ui selectable structured large table">
      <thead>
        <tr>
          <th colSpan="5">
            <h3>Selected tags</h3>
          </th>
        </tr>
      </thead>
      <tbody>
        {tagRows}
      </tbody>
    </table>
  );
}