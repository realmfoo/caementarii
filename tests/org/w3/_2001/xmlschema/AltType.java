//
// This file was generated by the JavaTM Architecture for XML Binding(JAXB) Reference Implementation, v2.2.8-b130911.1802 
// See <a href="http://java.sun.com/xml/jaxb">http://java.sun.com/xml/jaxb</a> 
// Any modifications to this file will be lost upon recompilation of the source schema. 
// Generated on: 2017.10.18 at 09:47:40 AM MSK 
//


package org.w3._2001.xmlschema;

import javax.xml.bind.annotation.XmlAccessType;
import javax.xml.bind.annotation.XmlAccessorType;
import javax.xml.bind.annotation.XmlAttribute;
import javax.xml.bind.annotation.XmlSchemaType;
import javax.xml.bind.annotation.XmlType;
import javax.xml.namespace.QName;


/**
 * 
 *         This type is used for 'alternative' elements.
 *       
 * 
 * <p>Java class for altType complex type.
 * 
 * <p>The following schema fragment specifies the expected content contained within this class.
 * 
 * <pre>
 * &lt;complexType name="altType">
 *   &lt;complexContent>
 *     &lt;extension base="{http://www.w3.org/2001/XMLSchema}annotated">
 *       &lt;choice minOccurs="0">
 *         &lt;element name="simpleType" type="{http://www.w3.org/2001/XMLSchema}localSimpleType"/>
 *         &lt;element name="complexType" type="{http://www.w3.org/2001/XMLSchema}localComplexType"/>
 *       &lt;/choice>
 *       &lt;attribute name="test" type="{http://www.w3.org/2001/XMLSchema}string" />
 *       &lt;attribute name="type" type="{http://www.w3.org/2001/XMLSchema}QName" />
 *       &lt;attribute name="xpathDefaultNamespace" type="{http://www.w3.org/2001/XMLSchema}xpathDefaultNamespace" />
 *       &lt;anyAttribute processContents='lax' namespace='##other'/>
 *     &lt;/extension>
 *   &lt;/complexContent>
 * &lt;/complexType>
 * </pre>
 * 
 * 
 */
@XmlAccessorType(XmlAccessType.FIELD)
@XmlType(name = "altType", propOrder = {
    "simpleType",
    "complexType"
})
public class AltType
    extends Annotated
{

    protected LocalSimpleType simpleType;
    protected LocalComplexType complexType;
    @XmlAttribute(name = "test")
    protected String test;
    @XmlAttribute(name = "type")
    protected QName type;
    @XmlAttribute(name = "xpathDefaultNamespace")
    @XmlSchemaType(name = "xpathDefaultNamespace")
    protected String xpathDefaultNamespace;

    /**
     * Gets the value of the simpleType property.
     * 
     * @return
     *     possible object is
     *     {@link LocalSimpleType }
     *     
     */
    public LocalSimpleType getSimpleType() {
        return simpleType;
    }

    /**
     * Sets the value of the simpleType property.
     * 
     * @param value
     *     allowed object is
     *     {@link LocalSimpleType }
     *     
     */
    public void setSimpleType(LocalSimpleType value) {
        this.simpleType = value;
    }

    /**
     * Gets the value of the complexType property.
     * 
     * @return
     *     possible object is
     *     {@link LocalComplexType }
     *     
     */
    public LocalComplexType getComplexType() {
        return complexType;
    }

    /**
     * Sets the value of the complexType property.
     * 
     * @param value
     *     allowed object is
     *     {@link LocalComplexType }
     *     
     */
    public void setComplexType(LocalComplexType value) {
        this.complexType = value;
    }

    /**
     * Gets the value of the test property.
     * 
     * @return
     *     possible object is
     *     {@link String }
     *     
     */
    public String getTest() {
        return test;
    }

    /**
     * Sets the value of the test property.
     * 
     * @param value
     *     allowed object is
     *     {@link String }
     *     
     */
    public void setTest(String value) {
        this.test = value;
    }

    /**
     * Gets the value of the type property.
     * 
     * @return
     *     possible object is
     *     {@link QName }
     *     
     */
    public QName getType() {
        return type;
    }

    /**
     * Sets the value of the type property.
     * 
     * @param value
     *     allowed object is
     *     {@link QName }
     *     
     */
    public void setType(QName value) {
        this.type = value;
    }

    /**
     * Gets the value of the xpathDefaultNamespace property.
     * 
     * @return
     *     possible object is
     *     {@link String }
     *     
     */
    public String getXpathDefaultNamespace() {
        return xpathDefaultNamespace;
    }

    /**
     * Sets the value of the xpathDefaultNamespace property.
     * 
     * @param value
     *     allowed object is
     *     {@link String }
     *     
     */
    public void setXpathDefaultNamespace(String value) {
        this.xpathDefaultNamespace = value;
    }

}
