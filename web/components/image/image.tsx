import "./image.css"

export function Image({ source: src }: { source: string }) {
  return <img className="image" src={src} />;
}
