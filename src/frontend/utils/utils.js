export function getYzmType(file) {
  function getYzmTypeBySize(widht, height) {
    switch ([widht, height]) {
      case [350, 80]:
        return 0;
      case [200, 60]:
        return 1;
      case [200, 80]:
        return 2;
      case [150, 45]:
        return 3;
    }
  }
  return new Promise((resolve, reject) => {
    const image = new Image();
    image.src = window.URL.createObjectURL(file);
    image.onload = () => {
      return getYzmTypeBySize(image.width, image.height);
    };
    image.onerror = reject;
  });
}