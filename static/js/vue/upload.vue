<template>
  <div>
    <h3>Upload</h3>
    <form enctype="multipart/form-data" action="upload" method="post">
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </div>
</template>