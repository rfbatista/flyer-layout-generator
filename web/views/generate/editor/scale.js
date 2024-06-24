export function calculateProportionalSize(
  originalWidth,
  originalHeight,
  containerWidth,
  containerHeight,
) {
  const originalAspectRatio = originalWidth / originalHeight;
  const containerAspectRatio = containerWidth / containerHeight;

  let newWidth, newHeight;

  if (originalAspectRatio > containerAspectRatio) {
    // Fit to container width
    newWidth = containerWidth;
    newHeight = containerWidth / originalAspectRatio;
  } else {
    // Fit to container height
    newWidth = containerHeight * originalAspectRatio;
    newHeight = containerHeight;
  }

  return {
    width: newWidth,
    height: newHeight,
    scale: newWidth / originalWidth,
  };
}

export function calScale(
  containerWidth,
  containerHeight,
  originalWidth,
  originalHeight,
) {
  const scaleX = containerWidth / originalWidth;
  const scaleY = containerHeight / originalHeight;
  const scale = Math.min(scaleX, scaleY);
  return scale;
}
