//
// This file was generated by the JavaTM Architecture for XML Binding(JAXB) Reference Implementation, v2.2.8-b130911.1802 
// See <a href="http://java.sun.com/xml/jaxb">http://java.sun.com/xml/jaxb</a> 
// Any modifications to this file will be lost upon recompilation of the source schema. 
// Generated on: 2017.10.18 at 09:47:40 AM MSK 
//


package org.w3._2001.xmlschema;

import javax.xml.bind.annotation.XmlEnum;
import javax.xml.bind.annotation.XmlEnumValue;
import javax.xml.bind.annotation.XmlType;


/**
 * <p>Java class for specialNamespaceList.
 * 
 * <p>The following schema fragment specifies the expected content contained within this class.
 * <p>
 * <pre>
 * &lt;simpleType name="specialNamespaceList">
 *   &lt;restriction base="{http://www.w3.org/2001/XMLSchema}token">
 *     &lt;enumeration value="##any"/>
 *     &lt;enumeration value="##other"/>
 *   &lt;/restriction>
 * &lt;/simpleType>
 * </pre>
 * 
 */
@XmlType(name = "specialNamespaceList")
@XmlEnum
public enum SpecialNamespaceList {

    @XmlEnumValue("##any")
    ANY("##any"),
    @XmlEnumValue("##other")
    OTHER("##other");
    private final String value;

    SpecialNamespaceList(String v) {
        value = v;
    }

    public String value() {
        return value;
    }

    public static SpecialNamespaceList fromValue(String v) {
        for (SpecialNamespaceList c: SpecialNamespaceList.values()) {
            if (c.value.equals(v)) {
                return c;
            }
        }
        throw new IllegalArgumentException(v);
    }

}
