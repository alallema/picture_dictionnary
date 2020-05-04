import React, { useState } from "react";
import { Grid, Image } from 'semantic-ui-react'

const formatPictureArray = ["JPEG", "JPG", "PNG", "PNG8", "PNG24", "GIF", "BMP", "WEBP", "RAW", "ICO", "PDF", "TIFF"];
const formatVideoArray = ["MOV", "MPEG4", "MP4", "AVI"];

export default function PictureResult(props) {
  let { pictures } = props;
  const [playing, togglePlayer] = useState(false);

  if (typeof pictures === "undefined"){
    pictures = []
  }
  
  const resultPictures = pictures.filter(picture => formatPictureArray.includes(picture.format.toUpperCase()));
  const resultVideos = pictures.filter(picture => formatVideoArray.includes(picture.format.toUpperCase()));

  const handleHover = (e) => {
    if (playing) {
      e.target.pause();
      e.target.fastSeek(0);
    }
    else
      e.target.play()
    togglePlayer(!playing)
  }

  const pictureRows = resultPictures.map((picture) => (
    <Grid.Column key={picture.id}>
        <Image
          src={picture.mainPictureURL}
        />
    </Grid.Column>
  ));
  const videoRows = resultVideos.map((video) => (
    <Grid.Column key={video.id}>
      <video controls width='250' muted onMouseOver={handleHover} onMouseOut={handleHover}>
        <source src={video.mainPictureURL}></source>
      </video>
    </Grid.Column>
  ));
  return (
    <Grid relaxed columns={4}>
        {pictureRows}
        {videoRows}
    </Grid>
  );
}