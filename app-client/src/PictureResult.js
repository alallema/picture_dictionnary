import React from "react";
import { Grid, Image } from 'semantic-ui-react'



export default function PictureResult(props) {
  let { pictures } = props;

  if (typeof pictures === "undefined"){
    pictures = []
  }
  const pictureRows = pictures.map((picture, idx) => (
    <Grid.Column key={picture.id}>
        <Image
          src={picture.mainPictureURL}
        />
    </Grid.Column>
  ));
  return (
    <Grid relaxed columns={4}>
        {pictureRows}
    </Grid>
  );
}