import "./loading_screen.css";
import GridLoader from "react-spinners/GridLoader";

// Optional: You can customize the spinner with CSS-in-JS

const LoadingScreen = ({ isLoading, children }: any) => {
  return (
    <div style={{ position: "relative" }}>
      {isLoading && (
        <div
          style={{
            position: "absolute",
            top: 0,
            left: 0,
            width: "100%",
            height: "100%",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            backgroundColor: "rgba(255, 255, 255, 0.8)",
            zIndex: 1000,
          }}
        >
          <GridLoader className="loading_screen__load-icon" size={20} />
        </div>
      )}
      <div style={{ visibility: isLoading ? "hidden" : "visible" }}>
        {children}
      </div>
    </div>
  );
};
export default LoadingScreen;
