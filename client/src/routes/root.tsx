import { Outlet, Link } from "react-router-dom";
import "./root.css"

const Root = () => {
    return (
        <>
            <nav className="sidebar">
                <div className="main-nav-container">
                    <Link to="/home" className="main-nav-link">Home</Link>
                    <Link to="/play/computer" className="main-nav-link">Play Computer</Link>
                    <Link to="/play/online" className="main-nav-link">Play Online</Link>
                </div>
                <Link to="/settings" className="minor-nav-link">Settings</Link>
            </nav>
            <div className="content-view">
                <Outlet />
            </div>
        </>
    );
};

export default Root;