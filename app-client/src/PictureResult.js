import React from "react";
import { Grid, Image, Embed } from 'semantic-ui-react'

const formatPictureArray = ["JPEG", "JPG", "PNG", "PNG8", "PNG24", "GIF", "BMP", "WEBP", "RAW", "ICO", "PDF", "TIFF"];
const formatVideoArray = ["MOV", "MPEG4", "MP4", "AVI"];

export default function PictureResult(props) {
  let { pictures } = props;

  if (typeof pictures === "undefined"){
    pictures = []
  }
  
  const resultPictures = pictures.filter(picture => formatPictureArray.includes(picture.format.toUpperCase()));
  const resultVideos = pictures.filter(picture => formatVideoArray.includes(picture.format.toUpperCase()));

  const pictureRows = resultPictures.map((picture) => (
    <Grid.Column key={picture.id}>
        <Image
          src={picture.mainPictureURL}
        />
    </Grid.Column>
  ));
  const videoRows = resultVideos.map((video) => (
    <Grid.Column key={video.id}>
    <Embed
      icon='right circle arrow'
      placeholder=''
      url={video.mainPictureURL}
    />
    </Grid.Column>
  ));
  return (
    <Grid relaxed columns={4}>
        {pictureRows}
        {videoRows}
    </Grid>
  );
}