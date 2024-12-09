'use client'
import React, { ChangeEvent, useRef } from 'react';

import styles from './styles.module.css';

const CreateProfile: React.FC = () => {
    const textRef = useRef<HTMLDivElement>(null);
    const [textColor, setTextColor] = React.useState<string>('#000');
    const [bgColor, setBgColor] = React.useState<string>('rgba(0,0,0,0)');

    const handleColorChange = (evt: ChangeEvent<HTMLInputElement>) => {
      if (textRef.current && evt) {
        const tag: string = {
          text: 'color',
          background: 'backgroundColor'
        }[evt.target.name] || 'color';
  
        
        textRef.current.style.setProperty(tag, evt.target.value);
        return  tag === 'color' 
            ? setTextColor(evt.target.value) 
            : setBgColor(evt.target.value);
      }
    };

  React.useEffect(() => {
    
    
  }, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Create Profile</h1>
      <form className={styles.form}>
        <label className={styles.label}>
          nickname:
          <input type="text" className={styles.input} maxLength={15} />
        </label>
        <label className={styles.labelColor}>
          text color:
          <span style={{backgroundColor: textColor}}/>
          <input 
          type="color" 
          className={styles.input} 
          name='text' 
          onChange={handleColorChange} 
          value={textColor}
          />
        </label>
        <label className={styles.labelColor}>
          background color:
          <span style={{backgroundColor: bgColor}}/>
          <input 
            type="color" 
            className={styles.input} 
            name='background' 
            onChange={handleColorChange} 
            value={bgColor}
            />
        </label>
        <div className={styles.result} ref={textRef}>
            meu texto ficara assim
        </div>
        <button type="submit" className={styles.button}>Create Profile</button>
      </form>
    </div>
  );
}

export default CreateProfile;