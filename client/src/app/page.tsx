import CreateProfile from "./CreateProfile";
import styles from "./page.module.css";

export default function Home() {

  return (
      <main className={styles.main}>
       <CreateProfile />
    </main>
  );
}
