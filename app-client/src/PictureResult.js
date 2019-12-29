import React from "react";
import { Grid, Image } from 'semantic-ui-react'

export default function PictureResult(props) {
  const { pictures } = props;

  console.log("PICTURE ROWS", pictures)
  const pictureRows = pictures.map((picture, idx) => (
    <Grid.Column>
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